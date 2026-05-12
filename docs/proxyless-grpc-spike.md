# Proxyless gRPC Spike — Initial Findings

## Context
Evaluating replacement of Istio sidecar proxy with native gRPC channel options.

## Results
- Memory: ~40% reduction per pod (no sidecar)
- Latency: p50 improved ~8ms, p99 improved ~12ms
- Trade-off: service discovery must move to DNS or xDS client

## Open questions
1. Load balancing strategy during transition
2. mTLS interaction with existing Istio setup
3. Rollback procedure

## Recommendation
Discuss in architecture review before committing direction.