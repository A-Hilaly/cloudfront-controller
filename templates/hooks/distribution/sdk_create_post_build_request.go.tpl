    // This is an idempotency token required in the API call...
    input.DistributionConfig.SetCallerReference(getIdempotencyToken())
    // This is because we can't have nice things...
	if input.DistributionConfig != nil {
		setQuantityFields(input.DistributionConfig)
	}
