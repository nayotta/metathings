/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.deviced.ConnectResponse');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ai.metathings.service.deviced.ErrorValue');
goog.require('proto.ai.metathings.service.deviced.StreamCallValue');
goog.require('proto.ai.metathings.service.deviced.UnaryCallValue');

goog.forwardDeclare('proto.ai.metathings.service.deviced.ConnectMessageKind');

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ai.metathings.service.deviced.ConnectResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.ai.metathings.service.deviced.ConnectResponse.oneofGroups_);
};
goog.inherits(proto.ai.metathings.service.deviced.ConnectResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.deviced.ConnectResponse.displayName = 'proto.ai.metathings.service.deviced.ConnectResponse';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.ai.metathings.service.deviced.ConnectResponse.oneofGroups_ = [[3,4,9]];

/**
 * @enum {number}
 */
proto.ai.metathings.service.deviced.ConnectResponse.UnionCase = {
  UNION_NOT_SET: 0,
  UNARY_CALL: 3,
  STREAM_CALL: 4,
  ERR: 9
};

/**
 * @return {proto.ai.metathings.service.deviced.ConnectResponse.UnionCase}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.getUnionCase = function() {
  return /** @type {proto.ai.metathings.service.deviced.ConnectResponse.UnionCase} */(jspb.Message.computeOneofCase(this, proto.ai.metathings.service.deviced.ConnectResponse.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.deviced.ConnectResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.deviced.ConnectResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.ConnectResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    sessionId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    kind: jspb.Message.getFieldWithDefault(msg, 2, 0),
    unaryCall: (f = msg.getUnaryCall()) && proto.ai.metathings.service.deviced.UnaryCallValue.toObject(includeInstance, f),
    streamCall: (f = msg.getStreamCall()) && proto.ai.metathings.service.deviced.StreamCallValue.toObject(includeInstance, f),
    err: (f = msg.getErr()) && proto.ai.metathings.service.deviced.ErrorValue.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ai.metathings.service.deviced.ConnectResponse}
 */
proto.ai.metathings.service.deviced.ConnectResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.deviced.ConnectResponse;
  return proto.ai.metathings.service.deviced.ConnectResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.deviced.ConnectResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.deviced.ConnectResponse}
 */
proto.ai.metathings.service.deviced.ConnectResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSessionId(value);
      break;
    case 2:
      var value = /** @type {!proto.ai.metathings.service.deviced.ConnectMessageKind} */ (reader.readEnum());
      msg.setKind(value);
      break;
    case 3:
      var value = new proto.ai.metathings.service.deviced.UnaryCallValue;
      reader.readMessage(value,proto.ai.metathings.service.deviced.UnaryCallValue.deserializeBinaryFromReader);
      msg.setUnaryCall(value);
      break;
    case 4:
      var value = new proto.ai.metathings.service.deviced.StreamCallValue;
      reader.readMessage(value,proto.ai.metathings.service.deviced.StreamCallValue.deserializeBinaryFromReader);
      msg.setStreamCall(value);
      break;
    case 9:
      var value = new proto.ai.metathings.service.deviced.ErrorValue;
      reader.readMessage(value,proto.ai.metathings.service.deviced.ErrorValue.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.deviced.ConnectResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.deviced.ConnectResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.ConnectResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSessionId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getKind();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getUnaryCall();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.ai.metathings.service.deviced.UnaryCallValue.serializeBinaryToWriter
    );
  }
  f = message.getStreamCall();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.ai.metathings.service.deviced.StreamCallValue.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      proto.ai.metathings.service.deviced.ErrorValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 session_id = 1;
 * @return {number}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.getSessionId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.setSessionId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional ConnectMessageKind kind = 2;
 * @return {!proto.ai.metathings.service.deviced.ConnectMessageKind}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.getKind = function() {
  return /** @type {!proto.ai.metathings.service.deviced.ConnectMessageKind} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.ai.metathings.service.deviced.ConnectMessageKind} value */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.setKind = function(value) {
  jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional UnaryCallValue unary_call = 3;
 * @return {?proto.ai.metathings.service.deviced.UnaryCallValue}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.getUnaryCall = function() {
  return /** @type{?proto.ai.metathings.service.deviced.UnaryCallValue} */ (
    jspb.Message.getWrapperField(this, proto.ai.metathings.service.deviced.UnaryCallValue, 3));
};


/** @param {?proto.ai.metathings.service.deviced.UnaryCallValue|undefined} value */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.setUnaryCall = function(value) {
  jspb.Message.setOneofWrapperField(this, 3, proto.ai.metathings.service.deviced.ConnectResponse.oneofGroups_[0], value);
};


proto.ai.metathings.service.deviced.ConnectResponse.prototype.clearUnaryCall = function() {
  this.setUnaryCall(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.hasUnaryCall = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional StreamCallValue stream_call = 4;
 * @return {?proto.ai.metathings.service.deviced.StreamCallValue}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.getStreamCall = function() {
  return /** @type{?proto.ai.metathings.service.deviced.StreamCallValue} */ (
    jspb.Message.getWrapperField(this, proto.ai.metathings.service.deviced.StreamCallValue, 4));
};


/** @param {?proto.ai.metathings.service.deviced.StreamCallValue|undefined} value */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.setStreamCall = function(value) {
  jspb.Message.setOneofWrapperField(this, 4, proto.ai.metathings.service.deviced.ConnectResponse.oneofGroups_[0], value);
};


proto.ai.metathings.service.deviced.ConnectResponse.prototype.clearStreamCall = function() {
  this.setStreamCall(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.hasStreamCall = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional ErrorValue err = 9;
 * @return {?proto.ai.metathings.service.deviced.ErrorValue}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.getErr = function() {
  return /** @type{?proto.ai.metathings.service.deviced.ErrorValue} */ (
    jspb.Message.getWrapperField(this, proto.ai.metathings.service.deviced.ErrorValue, 9));
};


/** @param {?proto.ai.metathings.service.deviced.ErrorValue|undefined} value */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.setErr = function(value) {
  jspb.Message.setOneofWrapperField(this, 9, proto.ai.metathings.service.deviced.ConnectResponse.oneofGroups_[0], value);
};


proto.ai.metathings.service.deviced.ConnectResponse.prototype.clearErr = function() {
  this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.ConnectResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 9) != null;
};

