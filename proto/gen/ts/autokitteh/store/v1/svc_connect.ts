// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file autokitteh/store/v1/svc.proto (package autokitteh.store.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DoRequest, DoResponse, GetRequest, GetResponse, ListRequest, ListResponse } from "./svc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service autokitteh.store.v1.StoreService
 */
export const StoreService = {
  typeName: "autokitteh.store.v1.StoreService",
  methods: {
    /**
     * @generated from rpc autokitteh.store.v1.StoreService.Do
     */
    do: {
      name: "Do",
      I: DoRequest,
      O: DoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.store.v1.StoreService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc autokitteh.store.v1.StoreService.List
     */
    list: {
      name: "List",
      I: ListRequest,
      O: ListResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

