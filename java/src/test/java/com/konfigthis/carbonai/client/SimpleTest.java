package com.konfigthis.carbonai.client;

import static org.junit.Assert.assertNotNull;

import org.junit.jupiter.api.Test;

public class SimpleTest {
    final String MOCK_SERVER_URL = "http://localhost:4010";

    @Test
    public void initClientTest() {
        Configuration configuration = new Configuration();
        configuration.host = "https://api.carbon.ai";
        
        configuration.accessToken  = "YOUR API KEY";
        
        configuration.apiKey  = "YOUR API KEY";
        
        configuration.customerId  = "YOUR API KEY";
        Carbon client = new Carbon(configuration);
        assertNotNull(client);
    }
}
