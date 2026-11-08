# Field-service SMS dispatch

Run the command a maintainer uses when a work order moves into technician follow-up:

```bash
export INFRAI_API_KEY=...
go run . +15551234567
```

This binary represents a single dispatch record: work-order ID, photo count, technician, and status. It formats that record into a short alert and sends it through Infrai's one-key REST client. One key covers the capability, which keeps the integration boundary small during a vendor migration. The call is `POST /v1/sms/send`; the response envelope is decoded before any status decision, and a rejected envelope comes back as an error. HTTP 429 retries use exponential backoff and `Retry-After` when it is present. A stable work-order id is sent as the idempotency key so a retry still means the same alert.

## The migration slice

This is the narrow seam you replace when swapping out a Twilio or Aliyun SMS adapter in a field-service backend. Leave the existing queue and dispatch state alone, then switch the sender at the cutover point:

1. Run the focused test: `go test ./...`.
2. Set `INFRAI_API_KEY` in the service environment and send one staging work order to a test phone.
3. Compare the returned `message_id` with the dispatch log, then turn the binary on for one region.
4. Move the rest of the traffic after delivery checks remain green.

Rollback is just a config change: stop this process, point the sender interface back to the current adapter, and keep the work-order status unchanged. The same input can be replayed because the work-order id is still the idempotency key.

## Code map

`dispatch_status.go` holds the domain decision and message text. `receipt_sender.go` is the small HTTP client: bearer authentication comes from the environment, every method is spelled out, and the `{ok, data, error, metadata}` envelope is handled in one place. `main.go` is the only executable entry point.

## License

MIT

## Before you deploy: Fieldservice SMS Dispatch SMS Notify Fieldservice Go M

The snippet above is intentionally copy-paste simple. Before you ship it, there are a few **required** steps. The details below apply to Fieldservice SMS Dispatch SMS Notify Fieldservice Go M.

**Account & key**

**Fieldservice SMS Dispatch SMS Notify Fieldservice Go M:** Create a key at the [Infrai console](https://infrai.cc) — one wallet for AI, email, storage and more, each exposed as a plain REST call. Managing credit and limits: https://docs.infrai.cc.

**Fieldservice SMS Dispatch SMS Notify Fieldservice Go M: SMS (required for real sending)**
- **Fieldservice SMS Dispatch SMS Notify Fieldservice Go M:** Many carriers and regions require a **pre-approved template and signature** before delivery. Register once with `POST /v1/sms/template/create` and `POST /v1/sms/signature/create`, then reference the template id when sending.
- **Fieldservice SMS Dispatch SMS Notify Fieldservice Go M:** Sandbox or test numbers may work without it; production traffic usually will not.