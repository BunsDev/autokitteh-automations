// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/sessions/v1/svc.proto (package autokitteh.sessions.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Session, SessionLog, SessionStateType } from "./session_pb.js";

/**
 * @generated from message autokitteh.sessions.v1.StartRequest
 */
export class StartRequest extends Message<StartRequest> {
  /**
   * @generated from field: autokitteh.sessions.v1.Session session = 1;
   */
  session?: Session;

  /**
   * Helper: if set, merged into and overwrites the session's inputs.
   *
   * @generated from field: map<string, string> json_inputs = 2;
   */
  jsonInputs: { [key: string]: string } = {};

  constructor(data?: PartialMessage<StartRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.StartRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session", kind: "message", T: Session },
    { no: 2, name: "json_inputs", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartRequest {
    return new StartRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartRequest {
    return new StartRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartRequest {
    return new StartRequest().fromJsonString(jsonString, options);
  }

  static equals(a: StartRequest | PlainMessage<StartRequest> | undefined, b: StartRequest | PlainMessage<StartRequest> | undefined): boolean {
    return proto3.util.equals(StartRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.StartResponse
 */
export class StartResponse extends Message<StartResponse> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId = "";

  constructor(data?: PartialMessage<StartResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.StartResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartResponse {
    return new StartResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartResponse {
    return new StartResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartResponse {
    return new StartResponse().fromJsonString(jsonString, options);
  }

  static equals(a: StartResponse | PlainMessage<StartResponse> | undefined, b: StartResponse | PlainMessage<StartResponse> | undefined): boolean {
    return proto3.util.equals(StartResponse, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.StopRequest
 */
export class StopRequest extends Message<StopRequest> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId = "";

  /**
   * @generated from field: string reason = 2;
   */
  reason = "";

  /**
   * non-graceful
   *
   * @generated from field: bool terminate = 3;
   */
  terminate = false;

  constructor(data?: PartialMessage<StopRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.StopRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "terminate", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StopRequest {
    return new StopRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StopRequest {
    return new StopRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StopRequest {
    return new StopRequest().fromJsonString(jsonString, options);
  }

  static equals(a: StopRequest | PlainMessage<StopRequest> | undefined, b: StopRequest | PlainMessage<StopRequest> | undefined): boolean {
    return proto3.util.equals(StopRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.StopResponse
 */
export class StopResponse extends Message<StopResponse> {
  constructor(data?: PartialMessage<StopResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.StopResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StopResponse {
    return new StopResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StopResponse {
    return new StopResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StopResponse {
    return new StopResponse().fromJsonString(jsonString, options);
  }

  static equals(a: StopResponse | PlainMessage<StopResponse> | undefined, b: StopResponse | PlainMessage<StopResponse> | undefined): boolean {
    return proto3.util.equals(StopResponse, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.ListRequest
 */
export class ListRequest extends Message<ListRequest> {
  /**
   * @generated from field: string deployment_id = 1;
   */
  deploymentId = "";

  /**
   * @generated from field: string env_id = 2;
   */
  envId = "";

  /**
   * @generated from field: string event_id = 3;
   */
  eventId = "";

  /**
   * @generated from field: string build_id = 4;
   */
  buildId = "";

  /**
   * @generated from field: autokitteh.sessions.v1.SessionStateType state_type = 5;
   */
  stateType = SessionStateType.UNSPECIFIED;

  /**
   * @generated from field: bool count_only = 10;
   */
  countOnly = false;

  /**
   * If the value is outside the allowed range, the sessions
   * gRPC service sets it to the closest range bound.
   *
   * @generated from field: int32 page_size = 20;
   */
  pageSize = 0;

  /**
   * @generated from field: int32 skip = 21;
   */
  skip = 0;

  /**
   * @generated from field: string page_token = 22;
   */
  pageToken = "";

  constructor(data?: PartialMessage<ListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.ListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "deployment_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "env_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "event_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "build_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "state_type", kind: "enum", T: proto3.getEnumType(SessionStateType) },
    { no: 10, name: "count_only", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 20, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 21, name: "skip", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 22, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRequest {
    return new ListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRequest {
    return new ListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRequest {
    return new ListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListRequest | PlainMessage<ListRequest> | undefined, b: ListRequest | PlainMessage<ListRequest> | undefined): boolean {
    return proto3.util.equals(ListRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.ListResponse
 */
export class ListResponse extends Message<ListResponse> {
  /**
   * Sessions without their data.
   *
   * @generated from field: repeated autokitteh.sessions.v1.Session sessions = 1;
   */
  sessions: Session[] = [];

  /**
   * @generated from field: int64 count = 2;
   */
  count = protoInt64.zero;

  /**
   * @generated from field: string next_page_token = 10;
   */
  nextPageToken = "";

  constructor(data?: PartialMessage<ListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.ListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sessions", kind: "message", T: Session, repeated: true },
    { no: 2, name: "count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 10, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListResponse {
    return new ListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListResponse {
    return new ListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListResponse {
    return new ListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListResponse | PlainMessage<ListResponse> | undefined, b: ListResponse | PlainMessage<ListResponse> | undefined): boolean {
    return proto3.util.equals(ListResponse, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.GetRequest
 */
export class GetRequest extends Message<GetRequest> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId = "";

  constructor(data?: PartialMessage<GetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.GetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRequest {
    return new GetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRequest {
    return new GetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRequest {
    return new GetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRequest | PlainMessage<GetRequest> | undefined, b: GetRequest | PlainMessage<GetRequest> | undefined): boolean {
    return proto3.util.equals(GetRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.GetResponse
 */
export class GetResponse extends Message<GetResponse> {
  /**
   * @generated from field: autokitteh.sessions.v1.Session session = 1;
   */
  session?: Session;

  constructor(data?: PartialMessage<GetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.GetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session", kind: "message", T: Session },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetResponse {
    return new GetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetResponse {
    return new GetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetResponse {
    return new GetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetResponse | PlainMessage<GetResponse> | undefined, b: GetResponse | PlainMessage<GetResponse> | undefined): boolean {
    return proto3.util.equals(GetResponse, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.GetLogRequest
 */
export class GetLogRequest extends Message<GetLogRequest> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId = "";

  /**
   * true: all values returned will be string values
   *       that contain the native values in JSON format.
   * false: all values returned are properly boxed.
   *
   * @generated from field: bool json_values = 2;
   */
  jsonValues = false;

  /**
   * @generated from field: bool ascending = 11;
   */
  ascending = false;

  /**
   * @generated from field: int32 page_size = 20;
   */
  pageSize = 0;

  /**
   * @generated from field: int32 skip = 21;
   */
  skip = 0;

  /**
   * @generated from field: string page_token = 22;
   */
  pageToken = "";

  constructor(data?: PartialMessage<GetLogRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.GetLogRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "json_values", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 11, name: "ascending", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 20, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 21, name: "skip", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 22, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLogRequest {
    return new GetLogRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLogRequest {
    return new GetLogRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLogRequest {
    return new GetLogRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetLogRequest | PlainMessage<GetLogRequest> | undefined, b: GetLogRequest | PlainMessage<GetLogRequest> | undefined): boolean {
    return proto3.util.equals(GetLogRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.GetLogResponse
 */
export class GetLogResponse extends Message<GetLogResponse> {
  /**
   * @generated from field: autokitteh.sessions.v1.SessionLog log = 1;
   */
  log?: SessionLog;

  /**
   * @generated from field: int64 count = 2;
   */
  count = protoInt64.zero;

  /**
   * @generated from field: string next_page_token = 10;
   */
  nextPageToken = "";

  constructor(data?: PartialMessage<GetLogResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.GetLogResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "log", kind: "message", T: SessionLog },
    { no: 2, name: "count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 10, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLogResponse {
    return new GetLogResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLogResponse {
    return new GetLogResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLogResponse {
    return new GetLogResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetLogResponse | PlainMessage<GetLogResponse> | undefined, b: GetLogResponse | PlainMessage<GetLogResponse> | undefined): boolean {
    return proto3.util.equals(GetLogResponse, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.DeleteRequest
 */
export class DeleteRequest extends Message<DeleteRequest> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId = "";

  constructor(data?: PartialMessage<DeleteRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.DeleteRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRequest {
    return new DeleteRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRequest {
    return new DeleteRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRequest {
    return new DeleteRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteRequest | PlainMessage<DeleteRequest> | undefined, b: DeleteRequest | PlainMessage<DeleteRequest> | undefined): boolean {
    return proto3.util.equals(DeleteRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.DeleteResponse
 */
export class DeleteResponse extends Message<DeleteResponse> {
  constructor(data?: PartialMessage<DeleteResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.DeleteResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteResponse {
    return new DeleteResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteResponse {
    return new DeleteResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteResponse {
    return new DeleteResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteResponse | PlainMessage<DeleteResponse> | undefined, b: DeleteResponse | PlainMessage<DeleteResponse> | undefined): boolean {
    return proto3.util.equals(DeleteResponse, a, b);
  }
}

