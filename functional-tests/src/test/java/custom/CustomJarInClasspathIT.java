/*
 * Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

package custom;

import static com.oracle.bedrock.deferred.DeferredHelper.invoking;
import static helm.HelmUtils.HELM_TIMEOUT;
import static helm.HelmUtils.getPods;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;

import java.io.File;
import java.net.URL;
import java.util.List;
import java.util.Queue;
import java.util.concurrent.TimeUnit;

import com.oracle.bedrock.deferred.options.MaximumRetryDelay;
import com.oracle.bedrock.deferred.options.RetryFrequency;
import com.oracle.bedrock.runtime.console.SystemApplicationConsole;
import com.tangosol.util.Resources;
import org.junit.After;
import org.junit.AfterClass;
import org.junit.BeforeClass;
import org.junit.Test;

import com.oracle.bedrock.options.Timeout;
import com.oracle.bedrock.runtime.k8s.K8sCluster;
import com.oracle.bedrock.runtime.k8s.helm.HelmUpgrade;
import com.oracle.bedrock.runtime.options.Arguments;
import com.oracle.bedrock.testsupport.deferred.Eventually;

import helm.BaseHelmChartTest;
import helm.HelmUtils;

/**
 * @author cp
 */
public class CustomJarInClasspathIT extends BaseHelmChartTest
    {
    // ----- test lifecycle --------------------------------------------------

    @BeforeClass
    public static void setup()
        {
        assertPreconditions(s_k8sCluster);
        ensureNamespace(s_k8sCluster);
        ensureSecret(s_k8sCluster);
        }

    @AfterClass
    public static void cleanup()
        {
        cleanupPullSecrets(s_k8sCluster);
        cleanupNamespace(s_k8sCluster);

        for (String sNamespace : getTargetNamespaces())
            {
            cleanupNamespace(s_k8sCluster, sNamespace);
            }
        }

    @After
    public void cleanUpCoherence()
        {
        if (m_sRelease != null)
            {
            deleteCoherence(s_k8sCluster, getK8sNamespace(), m_sRelease, PERSISTENCE);
            }
        }

    // ----- test methods ---------------------------------------------------

    /**
     * Run the test using supplied list of values files.
     *
     * @throws Exception  if the test fails
     */
    @Test
    public void testCoherenceWithUserSuppliedJarInClasspath() throws Exception
        {
        String sNamespace      = getK8sNamespace();
        String sValuesOriginal = "values/helm-values-coh-user-artifacts.yaml";
        String sValuesUpgrade  = "values/helm-values-coh-user-artifacts-upgrade.yaml";

        m_sRelease = installCoherence(s_k8sCluster,
                                      sNamespace,
                                      sValuesOriginal,
                                      "coherence.image=" + COHERENCE_IMAGE);

        assertCoherence(s_k8sCluster, sNamespace, m_sRelease);

        assertCoherenceService(s_k8sCluster, sNamespace, m_sRelease);

        String       sCoherenceSelector = getCoherencePodSelector(m_sRelease);
        List<String> listPods           = getPods(s_k8sCluster, sNamespace, sCoherenceSelector);

        assertThat(listPods.size(), is(3));

        System.err.println("Waiting for Coherence Pods");

        for (String sPod : listPods)
            {
            System.err.println("Waiting for Coherence Pod " + sPod + "...");
            Eventually.assertThat(invoking(this).hasDefaultCacheServerStarted(s_k8sCluster, sNamespace, sPod),
                                  is(true), Timeout.after(HELM_TIMEOUT, TimeUnit.SECONDS));
            }

        System.err.println("Coherence Pods started");

        try
            {
            installClient(s_k8sCluster, CLIENT1, sNamespace, m_sRelease, CLUSTER1);
            installClient(s_k8sCluster, CLIENT2, sNamespace, m_sRelease, CLUSTER1);

            System.err.println("Waiting for Client-1 initial state ...");
            Eventually.assertThat(invoking(this).isRequiredClientStateReached(s_k8sCluster, sNamespace, CLIENT1),
                                  is(true),
                                  MaximumRetryDelay.of(RETRY_FREQUENCEY_SECONDS, TimeUnit.SECONDS),
                                  RetryFrequency.every(RETRY_FREQUENCEY_SECONDS, TimeUnit.SECONDS),
                                  Eventually.within(TIMEOUT, TimeUnit.SECONDS));

            System.err.println("****************************************************************");
            System.err.println("********[STARTING UPGRADE OF " + m_sRelease + "]***********");
            System.err.println("****************************************************************");

            URL urlValues = Resources.findFileOrResource(sValuesUpgrade, null);

            upgradeUserArtifactsInCoherenceClasspath(m_sRelease, urlValues.getPath());

            System.err.println("Waiting for required Client-1 state after upgrade ...");
            Eventually.assertThat(invoking(this).isRequiredClientStateReachedAfterUpgrade(s_k8sCluster, sNamespace, CLIENT1),
                                  is(true),
                                  MaximumRetryDelay.of(RETRY_FREQUENCEY_SECONDS, TimeUnit.SECONDS),
                                  RetryFrequency.every(RETRY_FREQUENCEY_SECONDS, TimeUnit.SECONDS),
                                  Eventually.within(TIMEOUT, TimeUnit.SECONDS));
            }
        finally
            {
            dumpPodLog(s_k8sCluster, sNamespace, CLIENT1);
            deleteClients();
            }
        }

        /**
         * Run the test using supplied list of values files.
         *
         * @throws Exception  if the test fails
         */
        @Test
        public void testCoherenceServerWithUserSuppliedClass() throws Exception
        {
            String sNamespace      = getK8sNamespace();
            String sValuesOriginal = "values/helm-values-coh-user-artifacts-mainClass.yaml";

            m_sRelease = installCoherence(s_k8sCluster,
                    sNamespace,
                    sValuesOriginal,
                    "coherence.image=" + COHERENCE_IMAGE);

            assertCoherence(s_k8sCluster, sNamespace, m_sRelease);

            assertCoherenceService(s_k8sCluster, sNamespace, m_sRelease);

            String       sCoherenceSelector = getCoherencePodSelector(m_sRelease);
            List<String> listPods           = getPods(s_k8sCluster, sNamespace, sCoherenceSelector);

            assertThat(listPods.size(), is(3));

            System.err.println("Waiting for Coherence Pods");

            for (String sPod : listPods)
            {
                System.err.println("Waiting for Coherence Pod " + sPod + "...");
                Eventually.assertThat(invoking(this).hasDefaultCacheServerStarted(s_k8sCluster, sNamespace, sPod),
                        is(true), Timeout.after(HELM_TIMEOUT, TimeUnit.SECONDS));
            }

            System.err.println("Coherence Pods started");
        }

    // ----- helper methods -------------------------------------------------

    /**
     * Install coherence with upgraded user supplied artifacts
     *
     * @param sRelease     the release that is being upgraded to newer version of artifacts
     * @param sHelmValues  Helm values file
     */
    private void upgradeUserArtifactsInCoherenceClasspath(String sRelease, String sHelmValues) throws Exception
        {
        File fileTempDir = s_temp.newFolder();

        System.err.println("Extracting Helm chart " + COHERENCE_HELM_CHART_URL + " into " + fileTempDir);

        HelmUtils.extractTarGZ(fileTempDir, COHERENCE_HELM_CHART_URL);

        assertHelmLint(fileTempDir, COHERENCE_HELM_CHART_NAME);

        HelmUpgrade cohUpgrade = s_helm.upgrade(sRelease, fileTempDir + "/" + COHERENCE_HELM_CHART_NAME)
                                       .values(sHelmValues);

        if (K8S_IMAGE_PULL_SECRET != null && K8S_IMAGE_PULL_SECRET.trim().length() > 0)
            {
            cohUpgrade = cohUpgrade.set("imagePullSecrets={" + K8S_IMAGE_PULL_SECRET + "}");
            }

        int nExitCode = cohUpgrade.executeAndWait();

        if (nExitCode == 0)
            {
            captureInstalledPodLogs(getDefaultCluster(), getK8sNamespace(), sRelease);
            }

        assertThat("Helm upgrade returned non-zero exit code.", nExitCode, is(0));
        }

    /**
     * Check for required client state after the upgrade.
     *
     * @param cluster     the k8s cluster
     * @param sNamespace  the namespace name
     * @param sClientPod  the pod name
     *
     * @return {@code true} if required client state is reached.
     */
    // MUST BE PUBLIC - used in Eventually.assertThat
    public boolean isRequiredClientStateReachedAfterUpgrade(K8sCluster cluster, String sNamespace, String sClientPod)
        {
        try
            {
            Queue<String> sLogs = getPodLog(cluster, sNamespace, sClientPod, null, false);
            return sLogs.stream().anyMatch(l -> l.contains("Cache Value Before Cloud EntryProcessor: AWS"))
                    && sLogs.stream().anyMatch(l -> l.contains("Cache Value After Cloud EntryProcessor: OCI"));
            }
        catch (Exception ex)
            {
            return false;
            }
        }

    /**
     * Determine whether the coherence start up log is ready. param cluster the k8s
     * cluster
     *
     * @param cluster     the k8s cluster
     * @param sNamespace  the namespace of coherence
     * @param sPod        the pod name of coherence
     *
     * @return {@code true} if coherence container is ready
     */
    public boolean hasDefaultCacheServerStarted(K8sCluster cluster, String sNamespace, String sPod)
        {
        try
            {
            Queue<String> sLogs = getPodLog(cluster, sNamespace, sPod, COHERENCE_CONTAINER_NAME);
            return sLogs.stream().anyMatch(l -> l.contains("Started DefaultCacheServer"));
            }
        catch (Exception ex)
            {
            return false;
            }
        }

    private void deleteClients()
        {
        deleteClient(CLIENT1);
        deleteClient(CLIENT2);
        }

    private void deleteClient(String sClient)
        {
        String sNamespace = getK8sNamespace();
        Arguments arguments = Arguments.of("delete", "pod", sClient);

        if (sNamespace != null)
            {
            arguments = arguments.with("-n", sNamespace);
            }

        s_k8sCluster.kubectlAndWait(arguments, SystemApplicationConsole.builder());
        }

    // ----- data members ---------------------------------------------------

    /**
     * The k8s cluster to use to install the charts.
     */
    private static K8sCluster     s_k8sCluster = getDefaultCluster();

    /**
     * Time out value for checking the required condition.
     */
    private static final int      TIMEOUT      = 300;

    /**
     * The retry frequency in seconds.
     */
    private static final int RETRY_FREQUENCEY_SECONDS = 10;

    /**
     * The boolean indicates whether coherence cache data is persisted.
     */
    private static boolean        PERSISTENCE  = false;

    /**
     * The full Coherence image name to use.
     */
    public static final String COHERENCE_IMAGE = System.getProperty("test.coherence.image");

    /**
     * The name of the deployed Coherence Helm release.
     */
    private String                m_sRelease;

    private static final String   CLIENT1  = "coh-client-1";

    private static final String   CLIENT2  = "coh-client-2";

    private static final String   CLUSTER1 = "MyCoherenceCluster";
    }
