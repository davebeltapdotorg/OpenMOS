# MOS protocol source guide

Use the specification for the transport being changed. The MOS message names
overlap across generations, but their framing and envelopes do not.

| OpenMOS transport | Normative source | Boundary |
| --- | --- | --- |
| MOS 2.8.4 TCP | [MOS Protocol 2.8.4](https://mosprotocol.com/wp-content/MOS-Protocol-Documents/MOS-Protocol-2.8.4-Current.htm) | Raw TCP with UCS-2 big-endian XML. |
| MOS 4 WebSocket | [MOS Protocol 4.0](https://mosprotocol.com/wp-content/MOS-Protocol-Documents/MOS-Protocol-Version-4.0.pdf) | WebSocket binary frames with UCS-2 big-endian XML. |

The [MOS 3.8.4 WebService specification](https://mosprotocol.com/wp-content/MOS-Protocol-Documents/MOS-Protocol-3.8.4-Current.htm)
describes SOAP over HTTP; it does not set wire rules for either implemented
transport.

Apply these rules when changing the shared message handling:

- Keep framing and envelope validation in the owning transport. Accept an
  intentional compatibility shape on input without emitting it by default.
- Echo a request's `messageID` when a response is required; `keepAlive` has no
  response. A retry with the same ID and content must replay the original
  acknowledgement without applying the operation again; a reused ID with
  different content is a conflict.
- Acknowledge running-order writes only after successful storage. Preserve
  story and item order and carry `mosExternalMetadata` as opaque XML.
- Advertise only Profile 0. Parsing or handling selected running-order messages
  does not establish the complete Profile 2 workflow.

The [README](../README.md) states the implemented message scope. Local tests
establish local behavior; they do not prove interoperability with a live peer.
