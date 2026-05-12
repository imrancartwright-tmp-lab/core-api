func (w *BulkImportWorker) Run(ctx context.Context) error {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case job := <-w.queue:
            if err := w.process(ctx, job); err != nil {
                w.log.Error("job failed", "err", err)
            }
        }
    }
}
