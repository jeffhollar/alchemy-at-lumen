package com.lumen.workflow;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import io.temporal.client.WorkflowClient;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.serviceclient.WorkflowServiceStubsOptions;
import com.lumen.workflow.config.TemporalProperties;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SpringBootApplication
@EnableConfigurationProperties(TemporalProperties.class)
public class Application {
    private static final Logger logger = LoggerFactory.getLogger(Application.class);

    @Value("${security.api-key:#{null}}")
    private String apiKey;

    @Value("${security.ssl.server-root-ca-cert:#{null}}")
    private String serverRootCACert;

    @Value("${security.ssl.client-cert:#{null}}")
    private String clientCert;

    @Value("${security.ssl.client-key:#{null}}")
    private String clientKey;

    @Value("${security.ssl.server-name:#{null}}")
    private String serverName;

    @Value("${security.ssl.insecure-skip-verify:false}")
    private boolean insecureSkipVerify;

    private final TemporalProperties temporalProperties;

    public Application(TemporalProperties temporalProperties) {
        this.temporalProperties = temporalProperties;
    }

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Bean
    public WorkflowServiceStubs workflowServiceStubs() {
        String targetHost = temporalProperties.getTargetHost();
        logger.info("Configuring Temporal service with host: {}", targetHost);
        
        // Remove any leading or trailing whitespace and ensure proper format
        String cleanTargetHost = targetHost.trim();
        if (!cleanTargetHost.startsWith("dns:///")) {
            // If it doesn't start with dns:///, assume it's a direct host:port
            cleanTargetHost = cleanTargetHost.replaceAll("^dns://+", "");
            cleanTargetHost = cleanTargetHost.replaceAll("^//+", "");
        }
        
        logger.debug("Using cleaned target host: {}", cleanTargetHost);
        
        WorkflowServiceStubsOptions.Builder optionsBuilder = WorkflowServiceStubsOptions.newBuilder()
                .setTarget(cleanTargetHost);

        // Configure SSL if enabled
        if (clientCert != null && clientKey != null) {
            try {
                io.grpc.netty.shaded.io.netty.handler.ssl.SslContextBuilder sslBuilder = 
                    io.grpc.netty.shaded.io.netty.handler.ssl.SslContextBuilder.forClient();

                if (serverRootCACert != null) {
                    sslBuilder.trustManager(new java.io.File(serverRootCACert));
                }
                
                sslBuilder.keyManager(
                    new java.io.File(clientCert),
                    new java.io.File(clientKey)
                );
                
                optionsBuilder.setSslContext(sslBuilder.build())
                            .setEnableHttps(true);
                
                if (serverName != null) {
                    optionsBuilder.setTarget(serverName);
                }
            } catch (Exception e) {
                logger.error("Error configuring SSL", e);
                throw new RuntimeException("Failed to configure SSL", e);
            }
        }

        return WorkflowServiceStubs.newInstance(optionsBuilder.build());
    }

    @Bean
    public WorkflowClient workflowClient(WorkflowServiceStubs service) {
        String namespace = temporalProperties.getNamespace();
        logger.info("Creating Temporal workflow client for namespace: {}", namespace);
        return WorkflowClient.newInstance(service,
                io.temporal.client.WorkflowClientOptions.newBuilder()
                        .setNamespace(namespace)
                        .setIdentity(apiKey != null ? apiKey : "default-client")
                        .build());
    }
} 