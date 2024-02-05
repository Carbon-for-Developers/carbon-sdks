import { Carbon } from "./index";

describe("carbon-typescript-sdk", () => {
  it("initialize client", async () => {
    const carbon = new Carbon({
      apiKey: "API_KEY",
    });
  });
  const carbon = new Carbon({
    apiKey: "API_KEY",
    basePath: "http://127.0.0.1:4010",
  });
  it("get access token", async () => {
    const token = await carbon.auth.getAccessToken({
      customerId: "CUSTOMER_ID",
    });
    console.log(token.data);
  });
  it("get white labeling", async () => {
    const whiteLabeling = await carbon.auth.getWhiteLabeling();
    console.log(whiteLabeling.data);
  });
  it("query user data sources", async () => {
    const dataSources = await carbon.dataSources.queryUserDataSources({
      customerId: "CUSTOMER_ID",
    });
    console.log(dataSources.data);
  });
  it("sync s3 files", async () => {
    const syncS3Files = await carbon.integrations.syncS3Files({
      customerId: "CUSTOMER_ID",
      ids: [{ id: "ID", bucket: "BUCKET" }],
    });
    console.log(syncS3Files.data);
  });
  it("health check", async () => {
    const status = await carbon.health.check();
    console.log(status.data);
  });
});
