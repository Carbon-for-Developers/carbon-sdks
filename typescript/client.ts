/*
Carbon

Connect external data to LLMs, no matter the source.

The version of the OpenAPI document: 1.0.0


NOTE: This file is auto generated by Konfig (https://konfigthis.com).
*/

import { AxiosRequestConfig } from "axios";
import {
  AuthApi,
  CRMApi,
  DataSourcesApi,
  EmbeddingsApi,
  FilesApi,
  GithubApi,
  IntegrationsApi,
  OrganizationsApi,
  UsersApi,
  UtilitiesApi,
  WebhooksApi,
  WhiteLabelApi,
} from "./api";
import { Configuration, ConfigurationParameters } from "./configuration";
import { CarbonCustom } from "./client-custom";

export class Carbon extends CarbonCustom {
  readonly auth: AuthApi;
  readonly cRM: CRMApi;
  readonly dataSources: DataSourcesApi;
  readonly embeddings: EmbeddingsApi;
  readonly files: FilesApi;
  readonly github: GithubApi;
  readonly integrations: IntegrationsApi;
  readonly organizations: OrganizationsApi;
  readonly users: UsersApi;
  readonly utilities: UtilitiesApi;
  readonly webhooks: WebhooksApi;
  readonly whiteLabel: WhiteLabelApi;

  constructor(configurationParameters: ConfigurationParameters = {}) {
    super(configurationParameters);
    const configuration = new Configuration(configurationParameters);
    this.auth = new AuthApi(configuration);
    this.cRM = new CRMApi(configuration);
    this.dataSources = new DataSourcesApi(configuration);
    this.embeddings = new EmbeddingsApi(configuration);
    this.files = new FilesApi(configuration);
    this.github = new GithubApi(configuration);
    this.integrations = new IntegrationsApi(configuration);
    this.organizations = new OrganizationsApi(configuration);
    this.users = new UsersApi(configuration);
    this.utilities = new UtilitiesApi(configuration);
    this.webhooks = new WebhooksApi(configuration);
    this.whiteLabel = new WhiteLabelApi(configuration);
  }

}
