package com.lumen.workflow.options;

import io.temporal.client.WorkflowClientOptions;
import io.temporal.serviceclient.WorkflowServiceStubsOptions;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContextBuilder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.FileInputStream;
import java.io.ByteArrayInputStream;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.util.Properties;

public class ClientOptions {
    private static final Logger logger = LoggerFactory.getLogger(ClientOptions.class);
    
    private String targetHost = "192.168.1.233:7233";
    private String namespace = "nexus-lumen-act-ns";
    private String serverRootCACert;
    private String clientCert;
    private String clientKey;
    private String serverName;
    private boolean insecureSkipVerify = false;
    private String apiKey;

    public String getTargetHost() {
        return targetHost;
    }

    public String getNamespace() {
        return namespace;
    }
    
    public static ClientOptions fromProperties(String propertiesFile) {
        ClientOptions options = new ClientOptions();
        try {
            Properties props = new Properties();
            props.load(new FileInputStream(propertiesFile));
            
            options.targetHost = props.getProperty("target-host", options.targetHost);
            options.namespace = props.getProperty("namespace", options.namespace);
            options.serverRootCACert = props.getProperty("server-root-ca-cert");
            options.clientCert = props.getProperty("client-cert");
            options.clientKey = props.getProperty("client-key");
            options.serverName = props.getProperty("server-name");
            options.insecureSkipVerify = Boolean.parseBoolean(props.getProperty("insecure-skip-verify", "false"));
            options.apiKey = props.getProperty("api-key");
            
            validateOptions(options);
        } catch (Exception e) {
            logger.error("Error loading properties file", e);
            throw new RuntimeException("Failed to load client options", e);
        }
        return options;
    }
    
    private static void validateOptions(ClientOptions options) {
        if ((options.clientCert != null && options.clientKey == null) || 
            (options.clientCert == null && options.clientKey != null)) {
            throw new IllegalArgumentException("Either both or neither of client-key and client-cert are required");
        }
        if (options.clientCert != null && options.apiKey != null) {
            throw new IllegalArgumentException("Either client-cert and client-key or api-key are required, not both");
        }
    }
    
    public WorkflowServiceStubsOptions getServiceStubsOptions() {
        WorkflowServiceStubsOptions.Builder optionsBuilder = WorkflowServiceStubsOptions.newBuilder()
                .setTarget(targetHost);
        
        if (clientCert != null) {
            try {
                SslContext sslContext = buildSSLContext();
                optionsBuilder.setSslContext(sslContext);
            } catch (Exception e) {
                logger.error("Error building SSL context", e);
                throw new RuntimeException("Failed to build SSL context", e);
            }
        }
        
        return optionsBuilder.build();
    }
    
    public WorkflowClientOptions getClientOptions() {
        WorkflowClientOptions.Builder optionsBuilder = WorkflowClientOptions.newBuilder()
                .setNamespace(namespace);
        
        if (apiKey != null) {
            optionsBuilder.setIdentity(apiKey);
        }
        
        return optionsBuilder.build();
    }
    
    private SslContext buildSSLContext() throws Exception {
        SslContextBuilder builder = SslContextBuilder.forClient();
        
        // Load client certificate and key if provided
        if (clientCert != null && clientKey != null) {
            byte[] certBytes = Files.readAllBytes(Paths.get(clientCert));
            byte[] keyBytes = Files.readAllBytes(Paths.get(clientKey));
            
            ByteArrayInputStream certStream = new ByteArrayInputStream(certBytes);
            ByteArrayInputStream keyStream = new ByteArrayInputStream(keyBytes);
            
            builder.keyManager(certStream, keyStream);
        }
        
        // Load server CA if provided
        if (serverRootCACert != null) {
            byte[] caBytes = Files.readAllBytes(Paths.get(serverRootCACert));
            ByteArrayInputStream caStream = new ByteArrayInputStream(caBytes);
            
            builder.trustManager(caStream);
        }
        
        return builder.build();
    }
} 