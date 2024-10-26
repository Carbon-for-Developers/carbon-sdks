operation_parameter_map = {
    '/auth/v1/access_token-GET': {
        'parameters': [
        ]
    },
    '/auth/v1/white_labeling-GET': {
        'parameters': [
        ]
    },
    '/integrations/data/crm/accounts/{id}-GET': {
        'parameters': [
            {
                'name': 'id'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'includes'
            },
        ]
    },
    '/integrations/data/crm/accounts-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'includes'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
        ]
    },
    '/integrations/data/crm/contacts/{id}-GET': {
        'parameters': [
            {
                'name': 'id'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'includes'
            },
        ]
    },
    '/integrations/data/crm/contacts-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'includes'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
        ]
    },
    '/integrations/data/crm/leads/{id}-GET': {
        'parameters': [
            {
                'name': 'id'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'includes'
            },
        ]
    },
    '/integrations/data/crm/leads-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'includes'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
        ]
    },
    '/integrations/data/crm/opportunities-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'includes'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
        ]
    },
    '/integrations/data/crm/opportunities/{id}-GET': {
        'parameters': [
            {
                'name': 'id'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'includes'
            },
        ]
    },
    '/data_sources/tags/add-POST': {
        'parameters': [
            {
                'name': 'tags'
            },
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/user_data_sources-POST': {
        'parameters': [
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'filters'
            },
        ]
    },
    '/data_sources/tags/remove-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'tags_to_remove'
            },
            {
                'name': 'remove_all_tags'
            },
        ]
    },
    '/revoke_access_token-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/embeddings-POST': {
        'parameters': [
            {
                'name': 'query'
            },
            {
                'name': 'k'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'query_vector'
            },
            {
                'name': 'file_ids'
            },
            {
                'name': 'parent_file_ids'
            },
            {
                'name': 'include_all_children'
            },
            {
                'name': 'tags_v2'
            },
            {
                'name': 'include_tags'
            },
            {
                'name': 'include_vectors'
            },
            {
                'name': 'include_raw_file'
            },
            {
                'name': 'hybrid_search'
            },
            {
                'name': 'hybrid_search_tuning_parameters'
            },
            {
                'name': 'media_type'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'include_file_level_metadata'
            },
            {
                'name': 'high_accuracy'
            },
            {
                'name': 'rerank'
            },
            {
                'name': 'file_types_at_source'
            },
            {
                'name': 'exclude_cold_storage_files'
            },
        ]
    },
    '/text_chunks-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'include_vectors'
            },
        ]
    },
    '/list_chunks_and_embeddings-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'include_vectors'
            },
        ]
    },
    '/upload_chunks_and_embeddings-POST': {
        'parameters': [
            {
                'name': 'embedding_model'
            },
            {
                'name': 'chunks_and_embeddings'
            },
            {
                'name': 'overwrite_existing'
            },
            {
                'name': 'chunks_only'
            },
            {
                'name': 'custom_credentials'
            },
        ]
    },
    '/create_user_file_tags-POST': {
        'parameters': [
            {
                'name': 'tags'
            },
            {
                'name': 'organization_user_file_id'
            },
        ]
    },
    '/deletefile/{file_id}-DELETE': {
        'parameters': [
            {
                'name': 'file_id'
            },
        ]
    },
    '/delete_user_file_tags-POST': {
        'parameters': [
            {
                'name': 'tags'
            },
            {
                'name': 'organization_user_file_id'
            },
        ]
    },
    '/delete_files-POST': {
        'parameters': [
            {
                'name': 'file_ids'
            },
            {
                'name': 'sync_statuses'
            },
            {
                'name': 'delete_non_synced_only'
            },
            {
                'name': 'send_webhook'
            },
            {
                'name': 'delete_child_files'
            },
        ]
    },
    '/delete_files_v2-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'send_webhook'
            },
            {
                'name': 'preserve_file_record'
            },
        ]
    },
    '/parsed_file/{file_id}-GET': {
        'parameters': [
            {
                'name': 'file_id'
            },
        ]
    },
    '/raw_file/{file_id}-GET': {
        'parameters': [
            {
                'name': 'file_id'
            },
        ]
    },
    '/modify_cold_storage_parameters-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'enable_cold_storage'
            },
            {
                'name': 'hot_storage_time_to_live'
            },
        ]
    },
    '/move_to_hot_storage-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
        ]
    },
    '/user_files_v2-POST': {
        'parameters': [
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'include_raw_file'
            },
            {
                'name': 'include_parsed_text_file'
            },
            {
                'name': 'include_additional_files'
            },
            {
                'name': 'presigned_url_expiry_time_seconds'
            },
        ]
    },
    '/user_files-POST': {
        'parameters': [
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'include_raw_file'
            },
            {
                'name': 'include_parsed_text_file'
            },
            {
                'name': 'include_additional_files'
            },
            {
                'name': 'presigned_url_expiry_time_seconds'
            },
        ]
    },
    '/resync_file-POST': {
        'parameters': [
            {
                'name': 'file_id'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'force_embedding_generation'
            },
            {
                'name': 'skip_file_processing'
            },
        ]
    },
    '/uploadfile-POST': {
        'parameters': [
            {
                'name': 'file'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'use_ocr'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'detect_audio_language'
            },
            {
                'name': 'transcription_service'
            },
            {
                'name': 'include_speaker_labels'
            },
            {
                'name': 'media_type'
            },
            {
                'name': 'split_rows'
            },
            {
                'name': 'enable_cold_storage'
            },
            {
                'name': 'hot_storage_time_to_live'
            },
            {
                'name': 'generate_chunks_only'
            },
            {
                'name': 'store_file_only'
            },
        ]
    },
    '/upload_file_from_url-POST': {
        'parameters': [
            {
                'name': 'url'
            },
            {
                'name': 'file_name'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'use_textract'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'detect_audio_language'
            },
            {
                'name': 'transcription_service'
            },
            {
                'name': 'include_speaker_labels'
            },
            {
                'name': 'media_type'
            },
            {
                'name': 'split_rows'
            },
            {
                'name': 'cold_storage_params'
            },
            {
                'name': 'generate_chunks_only'
            },
            {
                'name': 'store_file_only'
            },
        ]
    },
    '/upload_text-POST': {
        'parameters': [
            {
                'name': 'contents'
            },
            {
                'name': 'name'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'overwrite_file_id'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'cold_storage_params'
            },
            {
                'name': 'generate_chunks_only'
            },
            {
                'name': 'store_file_only'
            },
        ]
    },
    '/integrations/data/github/issues/{issue_number}-GET': {
        'parameters': [
            {
                'name': 'issue_number'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
        ]
    },
    '/integrations/data/github/issues-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'page'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
        ]
    },
    '/integrations/data/github/pull_requests/{pull_number}-GET': {
        'parameters': [
            {
                'name': 'pull_number'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
        ]
    },
    '/integrations/data/github/pull_requests/comments-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
            {
                'name': 'pull_number'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'page'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
        ]
    },
    '/integrations/data/github/pull_requests/commits-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
            {
                'name': 'pull_number'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'page'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'next_cursor'
            },
        ]
    },
    '/integrations/data/github/pull_requests/files-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
            {
                'name': 'pull_number'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'page'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'next_cursor'
            },
        ]
    },
    '/integrations/data/github/pull_requests-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'repository'
            },
            {
                'name': 'include_remote_data'
            },
            {
                'name': 'page'
            },
            {
                'name': 'page_size'
            },
            {
                'name': 'next_cursor'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
        ]
    },
    '/integrations/items/sync/cancel-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/connect-POST': {
        'parameters': [
            {
                'name': 'authentication'
            },
            {
                'name': 'sync_options'
            },
        ]
    },
    '/integrations/freshdesk-POST': {
        'parameters': [
            {
                'name': 'domain'
            },
            {
                'name': 'api_key'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'sync_files_on_connection'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'file_sync_config'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/gitbook-POST': {
        'parameters': [
            {
                'name': 'organization'
            },
            {
                'name': 'access_token'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'sync_files_on_connection'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'file_sync_config'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/guru-POST': {
        'parameters': [
            {
                'name': 'username'
            },
            {
                'name': 'access_token'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'sync_files_on_connection'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'file_sync_config'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/s3-POST': {
        'parameters': [
            {
                'name': 'access_key'
            },
            {
                'name': 'access_key_secret'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'endpoint_url'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/oauth_url-POST': {
        'parameters': [
            {
                'name': 'service'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'scope'
            },
            {
                'name': 'scopes'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'zendesk_subdomain'
            },
            {
                'name': 'microsoft_tenant'
            },
            {
                'name': 'sharepoint_site_name'
            },
            {
                'name': 'confluence_subdomain'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'salesforce_domain'
            },
            {
                'name': 'sync_files_on_connection'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'connecting_new_account'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'use_ocr'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'enable_file_picker'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'incremental_sync'
            },
            {
                'name': 'file_sync_config'
            },
            {
                'name': 'automatically_open_file_picker'
            },
            {
                'name': 'gong_account_email'
            },
            {
                'name': 'servicenow_credentials'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/confluence/list-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'parent_id'
            },
        ]
    },
    '/integrations/slack/conversations-GET': {
        'parameters': [
            {
                'name': 'types'
            },
            {
                'name': 'cursor'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'exclude_archived'
            },
        ]
    },
    '/integrations/items/list-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'parent_id'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
        ]
    },
    '/integrations/outlook/user_folders-GET': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/gitbook/spaces-GET': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/gmail/user_labels-GET': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/outlook/user_categories-GET': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/github/repos-GET': {
        'parameters': [
            {
                'name': 'per_page'
            },
            {
                'name': 'page'
            },
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/azure_blob_storage/files-POST': {
        'parameters': [
            {
                'name': 'ids'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'use_ocr'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'file_sync_config'
            },
        ]
    },
    '/integrations/azure_blob_storage-POST': {
        'parameters': [
            {
                'name': 'account_name'
            },
            {
                'name': 'account_key'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/confluence/sync-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'ids'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'use_ocr'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'incremental_sync'
            },
            {
                'name': 'file_sync_config'
            },
        ]
    },
    '/integrations/items/sync-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/files/sync-POST': {
        'parameters': [
            {
                'name': 'data_source_id'
            },
            {
                'name': 'ids'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'use_ocr'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'incremental_sync'
            },
            {
                'name': 'file_sync_config'
            },
        ]
    },
    '/integrations/github-POST': {
        'parameters': [
            {
                'name': 'username'
            },
            {
                'name': 'access_token'
            },
            {
                'name': 'sync_source_items'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/gitbook/sync-POST': {
        'parameters': [
            {
                'name': 'space_ids'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'file_sync_config'
            },
        ]
    },
    '/integrations/gmail/sync-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'sync_attachments'
            },
            {
                'name': 'file_sync_config'
            },
            {
                'name': 'incremental_sync'
            },
        ]
    },
    '/integrations/outlook/sync-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'folder'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'sync_attachments'
            },
            {
                'name': 'file_sync_config'
            },
            {
                'name': 'incremental_sync'
            },
        ]
    },
    '/integrations/github/sync_repos-POST': {
        'parameters': [
            {
                'name': 'repos'
            },
            {
                'name': 'data_source_id'
            },
        ]
    },
    '/integrations/rss_feed-POST': {
        'parameters': [
            {
                'name': 'url'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'data_source_tags'
            },
        ]
    },
    '/integrations/s3/files-POST': {
        'parameters': [
            {
                'name': 'ids'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'max_items_per_chunk'
            },
            {
                'name': 'set_page_as_boundary'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'request_id'
            },
            {
                'name': 'use_ocr'
            },
            {
                'name': 'parse_pdf_tables_with_ocr'
            },
            {
                'name': 'file_sync_config'
            },
        ]
    },
    '/integrations/slack/sync-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'data_source_id'
            },
            {
                'name': 'request_id'
            },
        ]
    },
    '/organization-GET': {
        'parameters': [
        ]
    },
    '/organization/update-POST': {
        'parameters': [
            {
                'name': 'global_user_config'
            },
            {
                'name': 'data_source_configs'
            },
        ]
    },
    '/organization/statistics-POST': {
        'parameters': [
        ]
    },
    '/delete_users-POST': {
        'parameters': [
            {
                'name': 'customer_ids'
            },
        ]
    },
    '/user-POST': {
        'parameters': [
            {
                'name': 'customer_id'
            },
        ]
    },
    '/list_users-POST': {
        'parameters': [
            {
                'name': 'pagination'
            },
            {
                'name': 'filters'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'include_count'
            },
        ]
    },
    '/modify_user_configuration-POST': {
        'parameters': [
            {
                'name': 'configuration_key_name'
            },
            {
                'name': 'value'
            },
        ]
    },
    '/update_users-POST': {
        'parameters': [
            {
                'name': 'customer_ids'
            },
            {
                'name': 'auto_sync_enabled_sources'
            },
            {
                'name': 'max_files'
            },
            {
                'name': 'max_files_per_upload'
            },
            {
                'name': 'max_characters'
            },
            {
                'name': 'max_characters_per_file'
            },
            {
                'name': 'max_characters_per_upload'
            },
        ]
    },
    '/whoami-GET': {
        'parameters': [
        ]
    },
    '/fetch_urls-GET': {
        'parameters': [
            {
                'name': 'url'
            },
        ]
    },
    '/fetch_webpage-POST': {
        'parameters': [
            {
                'name': 'url'
            },
        ]
    },
    '/fetch_youtube_transcript-GET': {
        'parameters': [
            {
                'name': 'id'
            },
            {
                'name': 'raw'
            },
        ]
    },
    '/process_sitemap-GET': {
        'parameters': [
            {
                'name': 'url'
            },
        ]
    },
    '/scrape_sitemap-POST': {
        'parameters': [
            {
                'name': 'url'
            },
            {
                'name': 'tags'
            },
            {
                'name': 'max_pages_to_scrape'
            },
            {
                'name': 'chunk_size'
            },
            {
                'name': 'chunk_overlap'
            },
            {
                'name': 'skip_embedding_generation'
            },
            {
                'name': 'enable_auto_sync'
            },
            {
                'name': 'generate_sparse_vectors'
            },
            {
                'name': 'prepend_filename_to_chunks'
            },
            {
                'name': 'html_tags_to_skip'
            },
            {
                'name': 'css_classes_to_skip'
            },
            {
                'name': 'css_selectors_to_skip'
            },
            {
                'name': 'embedding_model'
            },
            {
                'name': 'url_paths_to_include'
            },
            {
                'name': 'url_paths_to_exclude'
            },
            {
                'name': 'urls_to_scrape'
            },
            {
                'name': 'download_css_and_media'
            },
            {
                'name': 'generate_chunks_only'
            },
            {
                'name': 'store_file_only'
            },
            {
                'name': 'use_premium_proxies'
            },
        ]
    },
    '/web_scrape-POST': {
        'parameters': [
        ]
    },
    '/search_urls-GET': {
        'parameters': [
            {
                'name': 'query'
            },
        ]
    },
    '/user_webpages-POST': {
        'parameters': [
            {
                'name': 'filters'
            },
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
        ]
    },
    '/add_webhook-POST': {
        'parameters': [
            {
                'name': 'url'
            },
        ]
    },
    '/delete_webhook/{webhook_id}-DELETE': {
        'parameters': [
            {
                'name': 'webhook_id'
            },
        ]
    },
    '/webhooks-POST': {
        'parameters': [
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'filters'
            },
        ]
    },
    '/white_label/create-POST': {
        'parameters': [
        ]
    },
    '/white_label/delete-POST': {
        'parameters': [
            {
                'name': 'ids'
            },
        ]
    },
    '/white_label/list-POST': {
        'parameters': [
            {
                'name': 'pagination'
            },
            {
                'name': 'order_by'
            },
            {
                'name': 'order_dir'
            },
            {
                'name': 'filters'
            },
        ]
    },
    '/white_label/update-POST': {
        'parameters': [
            {
                'name': 'data_source_type'
            },
            {
                'name': 'credentials'
            },
        ]
    },
};