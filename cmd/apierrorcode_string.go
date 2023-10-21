// Code generated by "stringer -type=APIErrorCode -trimprefix=Err api-errors.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrNone-0]
	_ = x[ErrAccessDenied-1]
	_ = x[ErrBadDigest-2]
	_ = x[ErrEntityTooSmall-3]
	_ = x[ErrEntityTooLarge-4]
	_ = x[ErrPolicyTooLarge-5]
	_ = x[ErrIncompleteBody-6]
	_ = x[ErrInternalError-7]
	_ = x[ErrInvalidAccessKeyID-8]
	_ = x[ErrAccessKeyDisabled-9]
	_ = x[ErrInvalidArgument-10]
	_ = x[ErrInvalidBucketName-11]
	_ = x[ErrInvalidDigest-12]
	_ = x[ErrInvalidRange-13]
	_ = x[ErrInvalidRangePartNumber-14]
	_ = x[ErrInvalidCopyPartRange-15]
	_ = x[ErrInvalidCopyPartRangeSource-16]
	_ = x[ErrInvalidMaxKeys-17]
	_ = x[ErrInvalidEncodingMethod-18]
	_ = x[ErrInvalidMaxUploads-19]
	_ = x[ErrInvalidMaxParts-20]
	_ = x[ErrInvalidPartNumberMarker-21]
	_ = x[ErrInvalidPartNumber-22]
	_ = x[ErrInvalidRequestBody-23]
	_ = x[ErrInvalidCopySource-24]
	_ = x[ErrInvalidMetadataDirective-25]
	_ = x[ErrInvalidCopyDest-26]
	_ = x[ErrInvalidPolicyDocument-27]
	_ = x[ErrInvalidObjectState-28]
	_ = x[ErrMalformedXML-29]
	_ = x[ErrMissingContentLength-30]
	_ = x[ErrMissingContentMD5-31]
	_ = x[ErrMissingRequestBodyError-32]
	_ = x[ErrMissingSecurityHeader-33]
	_ = x[ErrNoSuchBucket-34]
	_ = x[ErrNoSuchBucketPolicy-35]
	_ = x[ErrNoSuchBucketLifecycle-36]
	_ = x[ErrNoSuchLifecycleConfiguration-37]
	_ = x[ErrInvalidLifecycleWithObjectLock-38]
	_ = x[ErrNoSuchBucketSSEConfig-39]
	_ = x[ErrNoSuchCORSConfiguration-40]
	_ = x[ErrNoSuchWebsiteConfiguration-41]
	_ = x[ErrReplicationConfigurationNotFoundError-42]
	_ = x[ErrRemoteDestinationNotFoundError-43]
	_ = x[ErrReplicationDestinationMissingLock-44]
	_ = x[ErrRemoteTargetNotFoundError-45]
	_ = x[ErrReplicationRemoteConnectionError-46]
	_ = x[ErrReplicationBandwidthLimitError-47]
	_ = x[ErrBucketRemoteIdenticalToSource-48]
	_ = x[ErrBucketRemoteAlreadyExists-49]
	_ = x[ErrBucketRemoteLabelInUse-50]
	_ = x[ErrBucketRemoteArnTypeInvalid-51]
	_ = x[ErrBucketRemoteArnInvalid-52]
	_ = x[ErrBucketRemoteRemoveDisallowed-53]
	_ = x[ErrRemoteTargetNotVersionedError-54]
	_ = x[ErrReplicationSourceNotVersionedError-55]
	_ = x[ErrReplicationNeedsVersioningError-56]
	_ = x[ErrReplicationBucketNeedsVersioningError-57]
	_ = x[ErrReplicationDenyEditError-58]
	_ = x[ErrRemoteTargetDenyAddError-59]
	_ = x[ErrReplicationNoExistingObjects-60]
	_ = x[ErrReplicationValidationError-61]
	_ = x[ErrReplicationPermissionCheckError-62]
	_ = x[ErrObjectRestoreAlreadyInProgress-63]
	_ = x[ErrNoSuchKey-64]
	_ = x[ErrNoSuchUpload-65]
	_ = x[ErrInvalidVersionID-66]
	_ = x[ErrNoSuchVersion-67]
	_ = x[ErrNotImplemented-68]
	_ = x[ErrPreconditionFailed-69]
	_ = x[ErrRequestTimeTooSkewed-70]
	_ = x[ErrSignatureDoesNotMatch-71]
	_ = x[ErrMethodNotAllowed-72]
	_ = x[ErrInvalidPart-73]
	_ = x[ErrInvalidPartOrder-74]
	_ = x[ErrMissingPart-75]
	_ = x[ErrAuthorizationHeaderMalformed-76]
	_ = x[ErrMalformedPOSTRequest-77]
	_ = x[ErrPOSTFileRequired-78]
	_ = x[ErrSignatureVersionNotSupported-79]
	_ = x[ErrBucketNotEmpty-80]
	_ = x[ErrAllAccessDisabled-81]
	_ = x[ErrPolicyInvalidVersion-82]
	_ = x[ErrMissingFields-83]
	_ = x[ErrMissingCredTag-84]
	_ = x[ErrCredMalformed-85]
	_ = x[ErrInvalidRegion-86]
	_ = x[ErrInvalidServiceS3-87]
	_ = x[ErrInvalidServiceSTS-88]
	_ = x[ErrInvalidRequestVersion-89]
	_ = x[ErrMissingSignTag-90]
	_ = x[ErrMissingSignHeadersTag-91]
	_ = x[ErrMalformedDate-92]
	_ = x[ErrMalformedPresignedDate-93]
	_ = x[ErrMalformedCredentialDate-94]
	_ = x[ErrMalformedExpires-95]
	_ = x[ErrNegativeExpires-96]
	_ = x[ErrAuthHeaderEmpty-97]
	_ = x[ErrExpiredPresignRequest-98]
	_ = x[ErrRequestNotReadyYet-99]
	_ = x[ErrUnsignedHeaders-100]
	_ = x[ErrMissingDateHeader-101]
	_ = x[ErrInvalidQuerySignatureAlgo-102]
	_ = x[ErrInvalidQueryParams-103]
	_ = x[ErrBucketAlreadyOwnedByYou-104]
	_ = x[ErrInvalidDuration-105]
	_ = x[ErrBucketAlreadyExists-106]
	_ = x[ErrMetadataTooLarge-107]
	_ = x[ErrUnsupportedMetadata-108]
	_ = x[ErrUnsupportedHostHeader-109]
	_ = x[ErrMaximumExpires-110]
	_ = x[ErrSlowDownRead-111]
	_ = x[ErrSlowDownWrite-112]
	_ = x[ErrMaxVersionsExceeded-113]
	_ = x[ErrInvalidPrefixMarker-114]
	_ = x[ErrBadRequest-115]
	_ = x[ErrKeyTooLongError-116]
	_ = x[ErrInvalidBucketObjectLockConfiguration-117]
	_ = x[ErrObjectLockConfigurationNotFound-118]
	_ = x[ErrObjectLockConfigurationNotAllowed-119]
	_ = x[ErrNoSuchObjectLockConfiguration-120]
	_ = x[ErrObjectLocked-121]
	_ = x[ErrInvalidRetentionDate-122]
	_ = x[ErrPastObjectLockRetainDate-123]
	_ = x[ErrUnknownWORMModeDirective-124]
	_ = x[ErrBucketTaggingNotFound-125]
	_ = x[ErrObjectLockInvalidHeaders-126]
	_ = x[ErrInvalidTagDirective-127]
	_ = x[ErrPolicyAlreadyAttached-128]
	_ = x[ErrPolicyNotAttached-129]
	_ = x[ErrExcessData-130]
	_ = x[ErrIncompleteSignature-131]
	_ = x[ErrInvalidAction-132]
	_ = x[ErrInvalidClientTokenID-133]
	_ = x[ErrInvalidParameterCombination-134]
	_ = x[ErrInvalidQueryParameter-135]
	_ = x[ErrInvalidParameterValue-136]
	_ = x[ErrOptInRequired-137]
	_ = x[ErrRequestExpired-138]
	_ = x[ErrServiceUnavailable-139]
	_ = x[ErrThrottling-140]
	_ = x[ErrValidation-141]
	_ = x[ErrResourceNotFound-142]
	_ = x[ErrUnrecognizedClient-143]
	_ = x[ErrMalformedQueryString-144]
	_ = x[ErrSlowDown-145]
	_ = x[ErrInvalidSignature-146]
	_ = x[ErrRequestTimeout-147]
	_ = x[ErrNetworkDisconnected-148]
	_ = x[ErrUnknownError-149]
	_ = x[ErrObjectAlreadyInActiveTier-150]
	_ = x[ErrObjectNotInActiveTier-151]
	_ = x[ErrInvalidEncryptionMethod-152]
	_ = x[ErrInvalidEncryptionKeyID-153]
	_ = x[ErrInsecureSSECustomerRequest-154]
	_ = x[ErrSSEMultipartEncrypted-155]
	_ = x[ErrSSEEncryptedObject-156]
	_ = x[ErrInvalidEncryptionParameters-157]
	_ = x[ErrInvalidEncryptionParametersSSEC-158]
	_ = x[ErrInvalidSSECustomerAlgorithm-159]
	_ = x[ErrInvalidSSECustomerKey-160]
	_ = x[ErrMissingSSECustomerKey-161]
	_ = x[ErrMissingSSECustomerKeyMD5-162]
	_ = x[ErrSSECustomerKeyMD5Mismatch-163]
	_ = x[ErrInvalidSSECustomerParameters-164]
	_ = x[ErrIncompatibleEncryptionMethod-165]
	_ = x[ErrKMSNotConfigured-166]
	_ = x[ErrKMSKeyNotFoundException-167]
	_ = x[ErrKMSDefaultKeyAlreadyConfigured-168]
	_ = x[ErrNoAccessKey-169]
	_ = x[ErrInvalidToken-170]
	_ = x[ErrEventNotification-171]
	_ = x[ErrARNNotification-172]
	_ = x[ErrRegionNotification-173]
	_ = x[ErrOverlappingFilterNotification-174]
	_ = x[ErrFilterNameInvalid-175]
	_ = x[ErrFilterNamePrefix-176]
	_ = x[ErrFilterNameSuffix-177]
	_ = x[ErrFilterValueInvalid-178]
	_ = x[ErrOverlappingConfigs-179]
	_ = x[ErrUnsupportedNotification-180]
	_ = x[ErrContentSHA256Mismatch-181]
	_ = x[ErrContentChecksumMismatch-182]
	_ = x[ErrStorageFull-183]
	_ = x[ErrRequestBodyParse-184]
	_ = x[ErrObjectExistsAsDirectory-185]
	_ = x[ErrInvalidObjectName-186]
	_ = x[ErrInvalidObjectNamePrefixSlash-187]
	_ = x[ErrInvalidResourceName-188]
	_ = x[ErrInvalidLifecycleQueryParameter-189]
	_ = x[ErrServerNotInitialized-190]
	_ = x[ErrRequestTimedout-191]
	_ = x[ErrClientDisconnected-192]
	_ = x[ErrTooManyRequests-193]
	_ = x[ErrInvalidRequest-194]
	_ = x[ErrTransitionStorageClassNotFoundError-195]
	_ = x[ErrInvalidStorageClass-196]
	_ = x[ErrBackendDown-197]
	_ = x[ErrMalformedJSON-198]
	_ = x[ErrAdminNoSuchUser-199]
	_ = x[ErrAdminNoSuchUserLDAPWarn-200]
	_ = x[ErrAdminNoSuchGroup-201]
	_ = x[ErrAdminGroupNotEmpty-202]
	_ = x[ErrAdminGroupDisabled-203]
	_ = x[ErrAdminNoSuchJob-204]
	_ = x[ErrAdminNoSuchPolicy-205]
	_ = x[ErrAdminPolicyChangeAlreadyApplied-206]
	_ = x[ErrAdminInvalidArgument-207]
	_ = x[ErrAdminInvalidAccessKey-208]
	_ = x[ErrAdminInvalidSecretKey-209]
	_ = x[ErrAdminConfigNoQuorum-210]
	_ = x[ErrAdminConfigTooLarge-211]
	_ = x[ErrAdminConfigBadJSON-212]
	_ = x[ErrAdminNoSuchConfigTarget-213]
	_ = x[ErrAdminConfigEnvOverridden-214]
	_ = x[ErrAdminConfigDuplicateKeys-215]
	_ = x[ErrAdminConfigInvalidIDPType-216]
	_ = x[ErrAdminConfigLDAPNonDefaultConfigName-217]
	_ = x[ErrAdminConfigLDAPValidation-218]
	_ = x[ErrAdminConfigIDPCfgNameAlreadyExists-219]
	_ = x[ErrAdminConfigIDPCfgNameDoesNotExist-220]
	_ = x[ErrAdminCredentialsMismatch-221]
	_ = x[ErrInsecureClientRequest-222]
	_ = x[ErrObjectTampered-223]
	_ = x[ErrSiteReplicationInvalidRequest-224]
	_ = x[ErrSiteReplicationPeerResp-225]
	_ = x[ErrSiteReplicationBackendIssue-226]
	_ = x[ErrSiteReplicationServiceAccountError-227]
	_ = x[ErrSiteReplicationBucketConfigError-228]
	_ = x[ErrSiteReplicationBucketMetaError-229]
	_ = x[ErrSiteReplicationIAMError-230]
	_ = x[ErrSiteReplicationConfigMissing-231]
	_ = x[ErrAdminRebalanceAlreadyStarted-232]
	_ = x[ErrAdminRebalanceNotStarted-233]
	_ = x[ErrAdminBucketQuotaExceeded-234]
	_ = x[ErrAdminNoSuchQuotaConfiguration-235]
	_ = x[ErrHealNotImplemented-236]
	_ = x[ErrHealNoSuchProcess-237]
	_ = x[ErrHealInvalidClientToken-238]
	_ = x[ErrHealMissingBucket-239]
	_ = x[ErrHealAlreadyRunning-240]
	_ = x[ErrHealOverlappingPaths-241]
	_ = x[ErrIncorrectContinuationToken-242]
	_ = x[ErrEmptyRequestBody-243]
	_ = x[ErrUnsupportedFunction-244]
	_ = x[ErrInvalidExpressionType-245]
	_ = x[ErrBusy-246]
	_ = x[ErrUnauthorizedAccess-247]
	_ = x[ErrExpressionTooLong-248]
	_ = x[ErrIllegalSQLFunctionArgument-249]
	_ = x[ErrInvalidKeyPath-250]
	_ = x[ErrInvalidCompressionFormat-251]
	_ = x[ErrInvalidFileHeaderInfo-252]
	_ = x[ErrInvalidJSONType-253]
	_ = x[ErrInvalidQuoteFields-254]
	_ = x[ErrInvalidRequestParameter-255]
	_ = x[ErrInvalidDataType-256]
	_ = x[ErrInvalidTextEncoding-257]
	_ = x[ErrInvalidDataSource-258]
	_ = x[ErrInvalidTableAlias-259]
	_ = x[ErrMissingRequiredParameter-260]
	_ = x[ErrObjectSerializationConflict-261]
	_ = x[ErrUnsupportedSQLOperation-262]
	_ = x[ErrUnsupportedSQLStructure-263]
	_ = x[ErrUnsupportedSyntax-264]
	_ = x[ErrUnsupportedRangeHeader-265]
	_ = x[ErrLexerInvalidChar-266]
	_ = x[ErrLexerInvalidOperator-267]
	_ = x[ErrLexerInvalidLiteral-268]
	_ = x[ErrLexerInvalidIONLiteral-269]
	_ = x[ErrParseExpectedDatePart-270]
	_ = x[ErrParseExpectedKeyword-271]
	_ = x[ErrParseExpectedTokenType-272]
	_ = x[ErrParseExpected2TokenTypes-273]
	_ = x[ErrParseExpectedNumber-274]
	_ = x[ErrParseExpectedRightParenBuiltinFunctionCall-275]
	_ = x[ErrParseExpectedTypeName-276]
	_ = x[ErrParseExpectedWhenClause-277]
	_ = x[ErrParseUnsupportedToken-278]
	_ = x[ErrParseUnsupportedLiteralsGroupBy-279]
	_ = x[ErrParseExpectedMember-280]
	_ = x[ErrParseUnsupportedSelect-281]
	_ = x[ErrParseUnsupportedCase-282]
	_ = x[ErrParseUnsupportedCaseClause-283]
	_ = x[ErrParseUnsupportedAlias-284]
	_ = x[ErrParseUnsupportedSyntax-285]
	_ = x[ErrParseUnknownOperator-286]
	_ = x[ErrParseMissingIdentAfterAt-287]
	_ = x[ErrParseUnexpectedOperator-288]
	_ = x[ErrParseUnexpectedTerm-289]
	_ = x[ErrParseUnexpectedToken-290]
	_ = x[ErrParseUnexpectedKeyword-291]
	_ = x[ErrParseExpectedExpression-292]
	_ = x[ErrParseExpectedLeftParenAfterCast-293]
	_ = x[ErrParseExpectedLeftParenValueConstructor-294]
	_ = x[ErrParseExpectedLeftParenBuiltinFunctionCall-295]
	_ = x[ErrParseExpectedArgumentDelimiter-296]
	_ = x[ErrParseCastArity-297]
	_ = x[ErrParseInvalidTypeParam-298]
	_ = x[ErrParseEmptySelect-299]
	_ = x[ErrParseSelectMissingFrom-300]
	_ = x[ErrParseExpectedIdentForGroupName-301]
	_ = x[ErrParseExpectedIdentForAlias-302]
	_ = x[ErrParseUnsupportedCallWithStar-303]
	_ = x[ErrParseNonUnaryAgregateFunctionCall-304]
	_ = x[ErrParseMalformedJoin-305]
	_ = x[ErrParseExpectedIdentForAt-306]
	_ = x[ErrParseAsteriskIsNotAloneInSelectList-307]
	_ = x[ErrParseCannotMixSqbAndWildcardInSelectList-308]
	_ = x[ErrParseInvalidContextForWildcardInSelectList-309]
	_ = x[ErrIncorrectSQLFunctionArgumentType-310]
	_ = x[ErrValueParseFailure-311]
	_ = x[ErrEvaluatorInvalidArguments-312]
	_ = x[ErrIntegerOverflow-313]
	_ = x[ErrLikeInvalidInputs-314]
	_ = x[ErrCastFailed-315]
	_ = x[ErrInvalidCast-316]
	_ = x[ErrEvaluatorInvalidTimestampFormatPattern-317]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbolForParsing-318]
	_ = x[ErrEvaluatorTimestampFormatPatternDuplicateFields-319]
	_ = x[ErrEvaluatorTimestampFormatPatternHourClockAmPmMismatch-320]
	_ = x[ErrEvaluatorUnterminatedTimestampFormatPatternToken-321]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternToken-322]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbol-323]
	_ = x[ErrEvaluatorBindingDoesNotExist-324]
	_ = x[ErrMissingHeaders-325]
	_ = x[ErrInvalidColumnIndex-326]
	_ = x[ErrAdminConfigNotificationTargetsFailed-327]
	_ = x[ErrAdminProfilerNotEnabled-328]
	_ = x[ErrInvalidDecompressedSize-329]
	_ = x[ErrAddUserInvalidArgument-330]
	_ = x[ErrAdminResourceInvalidArgument-331]
	_ = x[ErrAdminAccountNotEligible-332]
	_ = x[ErrAccountNotEligible-333]
	_ = x[ErrAdminServiceAccountNotFound-334]
	_ = x[ErrPostPolicyConditionInvalidFormat-335]
	_ = x[ErrInvalidChecksum-336]
	_ = x[ErrLambdaARNInvalid-337]
	_ = x[ErrLambdaARNNotFound-338]
	_ = x[apiErrCodeEnd-339]
}

const _APIErrorCode_name = "NoneAccessDeniedBadDigestEntityTooSmallEntityTooLargePolicyTooLargeIncompleteBodyInternalErrorInvalidAccessKeyIDAccessKeyDisabledInvalidArgumentInvalidBucketNameInvalidDigestInvalidRangeInvalidRangePartNumberInvalidCopyPartRangeInvalidCopyPartRangeSourceInvalidMaxKeysInvalidEncodingMethodInvalidMaxUploadsInvalidMaxPartsInvalidPartNumberMarkerInvalidPartNumberInvalidRequestBodyInvalidCopySourceInvalidMetadataDirectiveInvalidCopyDestInvalidPolicyDocumentInvalidObjectStateMalformedXMLMissingContentLengthMissingContentMD5MissingRequestBodyErrorMissingSecurityHeaderNoSuchBucketNoSuchBucketPolicyNoSuchBucketLifecycleNoSuchLifecycleConfigurationInvalidLifecycleWithObjectLockNoSuchBucketSSEConfigNoSuchCORSConfigurationNoSuchWebsiteConfigurationReplicationConfigurationNotFoundErrorRemoteDestinationNotFoundErrorReplicationDestinationMissingLockRemoteTargetNotFoundErrorReplicationRemoteConnectionErrorReplicationBandwidthLimitErrorBucketRemoteIdenticalToSourceBucketRemoteAlreadyExistsBucketRemoteLabelInUseBucketRemoteArnTypeInvalidBucketRemoteArnInvalidBucketRemoteRemoveDisallowedRemoteTargetNotVersionedErrorReplicationSourceNotVersionedErrorReplicationNeedsVersioningErrorReplicationBucketNeedsVersioningErrorReplicationDenyEditErrorRemoteTargetDenyAddErrorReplicationNoExistingObjectsReplicationValidationErrorReplicationPermissionCheckErrorObjectRestoreAlreadyInProgressNoSuchKeyNoSuchUploadInvalidVersionIDNoSuchVersionNotImplementedPreconditionFailedRequestTimeTooSkewedSignatureDoesNotMatchMethodNotAllowedInvalidPartInvalidPartOrderMissingPartAuthorizationHeaderMalformedMalformedPOSTRequestPOSTFileRequiredSignatureVersionNotSupportedBucketNotEmptyAllAccessDisabledPolicyInvalidVersionMissingFieldsMissingCredTagCredMalformedInvalidRegionInvalidServiceS3InvalidServiceSTSInvalidRequestVersionMissingSignTagMissingSignHeadersTagMalformedDateMalformedPresignedDateMalformedCredentialDateMalformedExpiresNegativeExpiresAuthHeaderEmptyExpiredPresignRequestRequestNotReadyYetUnsignedHeadersMissingDateHeaderInvalidQuerySignatureAlgoInvalidQueryParamsBucketAlreadyOwnedByYouInvalidDurationBucketAlreadyExistsMetadataTooLargeUnsupportedMetadataUnsupportedHostHeaderMaximumExpiresSlowDownReadSlowDownWriteMaxVersionsExceededInvalidPrefixMarkerBadRequestKeyTooLongErrorInvalidBucketObjectLockConfigurationObjectLockConfigurationNotFoundObjectLockConfigurationNotAllowedNoSuchObjectLockConfigurationObjectLockedInvalidRetentionDatePastObjectLockRetainDateUnknownWORMModeDirectiveBucketTaggingNotFoundObjectLockInvalidHeadersInvalidTagDirectivePolicyAlreadyAttachedPolicyNotAttachedExcessDataIncompleteSignatureInvalidActionInvalidClientTokenIDInvalidParameterCombinationInvalidQueryParameterInvalidParameterValueOptInRequiredRequestExpiredServiceUnavailableThrottlingValidationResourceNotFoundUnrecognizedClientMalformedQueryStringSlowDownInvalidSignatureRequestTimeoutNetworkDisconnectedUnknownErrorObjectAlreadyInActiveTierObjectNotInActiveTierInvalidEncryptionMethodInvalidEncryptionKeyIDInsecureSSECustomerRequestSSEMultipartEncryptedSSEEncryptedObjectInvalidEncryptionParametersInvalidEncryptionParametersSSECInvalidSSECustomerAlgorithmInvalidSSECustomerKeyMissingSSECustomerKeyMissingSSECustomerKeyMD5SSECustomerKeyMD5MismatchInvalidSSECustomerParametersIncompatibleEncryptionMethodKMSNotConfiguredKMSKeyNotFoundExceptionKMSDefaultKeyAlreadyConfiguredNoAccessKeyInvalidTokenEventNotificationARNNotificationRegionNotificationOverlappingFilterNotificationFilterNameInvalidFilterNamePrefixFilterNameSuffixFilterValueInvalidOverlappingConfigsUnsupportedNotificationContentSHA256MismatchContentChecksumMismatchStorageFullRequestBodyParseObjectExistsAsDirectoryInvalidObjectNameInvalidObjectNamePrefixSlashInvalidResourceNameInvalidLifecycleQueryParameterServerNotInitializedRequestTimedoutClientDisconnectedTooManyRequestsInvalidRequestTransitionStorageClassNotFoundErrorInvalidStorageClassBackendDownMalformedJSONAdminNoSuchUserAdminNoSuchUserLDAPWarnAdminNoSuchGroupAdminGroupNotEmptyAdminGroupDisabledAdminNoSuchJobAdminNoSuchPolicyAdminPolicyChangeAlreadyAppliedAdminInvalidArgumentAdminInvalidAccessKeyAdminInvalidSecretKeyAdminConfigNoQuorumAdminConfigTooLargeAdminConfigBadJSONAdminNoSuchConfigTargetAdminConfigEnvOverriddenAdminConfigDuplicateKeysAdminConfigInvalidIDPTypeAdminConfigLDAPNonDefaultConfigNameAdminConfigLDAPValidationAdminConfigIDPCfgNameAlreadyExistsAdminConfigIDPCfgNameDoesNotExistAdminCredentialsMismatchInsecureClientRequestObjectTamperedSiteReplicationInvalidRequestSiteReplicationPeerRespSiteReplicationBackendIssueSiteReplicationServiceAccountErrorSiteReplicationBucketConfigErrorSiteReplicationBucketMetaErrorSiteReplicationIAMErrorSiteReplicationConfigMissingAdminRebalanceAlreadyStartedAdminRebalanceNotStartedAdminBucketQuotaExceededAdminNoSuchQuotaConfigurationHealNotImplementedHealNoSuchProcessHealInvalidClientTokenHealMissingBucketHealAlreadyRunningHealOverlappingPathsIncorrectContinuationTokenEmptyRequestBodyUnsupportedFunctionInvalidExpressionTypeBusyUnauthorizedAccessExpressionTooLongIllegalSQLFunctionArgumentInvalidKeyPathInvalidCompressionFormatInvalidFileHeaderInfoInvalidJSONTypeInvalidQuoteFieldsInvalidRequestParameterInvalidDataTypeInvalidTextEncodingInvalidDataSourceInvalidTableAliasMissingRequiredParameterObjectSerializationConflictUnsupportedSQLOperationUnsupportedSQLStructureUnsupportedSyntaxUnsupportedRangeHeaderLexerInvalidCharLexerInvalidOperatorLexerInvalidLiteralLexerInvalidIONLiteralParseExpectedDatePartParseExpectedKeywordParseExpectedTokenTypeParseExpected2TokenTypesParseExpectedNumberParseExpectedRightParenBuiltinFunctionCallParseExpectedTypeNameParseExpectedWhenClauseParseUnsupportedTokenParseUnsupportedLiteralsGroupByParseExpectedMemberParseUnsupportedSelectParseUnsupportedCaseParseUnsupportedCaseClauseParseUnsupportedAliasParseUnsupportedSyntaxParseUnknownOperatorParseMissingIdentAfterAtParseUnexpectedOperatorParseUnexpectedTermParseUnexpectedTokenParseUnexpectedKeywordParseExpectedExpressionParseExpectedLeftParenAfterCastParseExpectedLeftParenValueConstructorParseExpectedLeftParenBuiltinFunctionCallParseExpectedArgumentDelimiterParseCastArityParseInvalidTypeParamParseEmptySelectParseSelectMissingFromParseExpectedIdentForGroupNameParseExpectedIdentForAliasParseUnsupportedCallWithStarParseNonUnaryAgregateFunctionCallParseMalformedJoinParseExpectedIdentForAtParseAsteriskIsNotAloneInSelectListParseCannotMixSqbAndWildcardInSelectListParseInvalidContextForWildcardInSelectListIncorrectSQLFunctionArgumentTypeValueParseFailureEvaluatorInvalidArgumentsIntegerOverflowLikeInvalidInputsCastFailedInvalidCastEvaluatorInvalidTimestampFormatPatternEvaluatorInvalidTimestampFormatPatternSymbolForParsingEvaluatorTimestampFormatPatternDuplicateFieldsEvaluatorTimestampFormatPatternHourClockAmPmMismatchEvaluatorUnterminatedTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternSymbolEvaluatorBindingDoesNotExistMissingHeadersInvalidColumnIndexAdminConfigNotificationTargetsFailedAdminProfilerNotEnabledInvalidDecompressedSizeAddUserInvalidArgumentAdminResourceInvalidArgumentAdminAccountNotEligibleAccountNotEligibleAdminServiceAccountNotFoundPostPolicyConditionInvalidFormatInvalidChecksumLambdaARNInvalidLambdaARNNotFoundapiErrCodeEnd"

var _APIErrorCode_index = [...]uint16{0, 4, 16, 25, 39, 53, 67, 81, 94, 112, 129, 144, 161, 174, 186, 208, 228, 254, 268, 289, 306, 321, 344, 361, 379, 396, 420, 435, 456, 474, 486, 506, 523, 546, 567, 579, 597, 618, 646, 676, 697, 720, 746, 783, 813, 846, 871, 903, 933, 962, 987, 1009, 1035, 1057, 1085, 1114, 1148, 1179, 1216, 1240, 1264, 1292, 1318, 1349, 1379, 1388, 1400, 1416, 1429, 1443, 1461, 1481, 1502, 1518, 1529, 1545, 1556, 1584, 1604, 1620, 1648, 1662, 1679, 1699, 1712, 1726, 1739, 1752, 1768, 1785, 1806, 1820, 1841, 1854, 1876, 1899, 1915, 1930, 1945, 1966, 1984, 1999, 2016, 2041, 2059, 2082, 2097, 2116, 2132, 2151, 2172, 2186, 2198, 2211, 2230, 2249, 2259, 2274, 2310, 2341, 2374, 2403, 2415, 2435, 2459, 2483, 2504, 2528, 2547, 2568, 2585, 2595, 2614, 2627, 2647, 2674, 2695, 2716, 2729, 2743, 2761, 2771, 2781, 2797, 2815, 2835, 2843, 2859, 2873, 2892, 2904, 2929, 2950, 2973, 2995, 3021, 3042, 3060, 3087, 3118, 3145, 3166, 3187, 3211, 3236, 3264, 3292, 3308, 3331, 3361, 3372, 3384, 3401, 3416, 3434, 3463, 3480, 3496, 3512, 3530, 3548, 3571, 3592, 3615, 3626, 3642, 3665, 3682, 3710, 3729, 3759, 3779, 3794, 3812, 3827, 3841, 3876, 3895, 3906, 3919, 3934, 3957, 3973, 3991, 4009, 4023, 4040, 4071, 4091, 4112, 4133, 4152, 4171, 4189, 4212, 4236, 4260, 4285, 4320, 4345, 4379, 4412, 4436, 4457, 4471, 4500, 4523, 4550, 4584, 4616, 4646, 4669, 4697, 4725, 4749, 4773, 4802, 4820, 4837, 4859, 4876, 4894, 4914, 4940, 4956, 4975, 4996, 5000, 5018, 5035, 5061, 5075, 5099, 5120, 5135, 5153, 5176, 5191, 5210, 5227, 5244, 5268, 5295, 5318, 5341, 5358, 5380, 5396, 5416, 5435, 5457, 5478, 5498, 5520, 5544, 5563, 5605, 5626, 5649, 5670, 5701, 5720, 5742, 5762, 5788, 5809, 5831, 5851, 5875, 5898, 5917, 5937, 5959, 5982, 6013, 6051, 6092, 6122, 6136, 6157, 6173, 6195, 6225, 6251, 6279, 6312, 6330, 6353, 6388, 6428, 6470, 6502, 6519, 6544, 6559, 6576, 6586, 6597, 6635, 6689, 6735, 6787, 6835, 6878, 6922, 6950, 6964, 6982, 7018, 7041, 7064, 7086, 7114, 7137, 7155, 7182, 7214, 7229, 7245, 7262, 7275}

func (i APIErrorCode) String() string {
	if i < 0 || i >= APIErrorCode(len(_APIErrorCode_index)-1) {
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APIErrorCode_name[_APIErrorCode_index[i]:_APIErrorCode_index[i+1]]
}
