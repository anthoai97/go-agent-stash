package business

func (b *Business) SyncToS3(from, to string) error {
	err := b.s3Storage.Sync(from, to)
	return err
}

// Daily ?
