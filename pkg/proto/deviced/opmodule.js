/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.deviced.OpModule');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.google.protobuf.StringValue');
goog.require('proto.google.protobuf.Timestamp');

goog.forwardDeclare('proto.ai.metathings.constant.state.ModuleState');

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
proto.ai.metathings.service.deviced.OpModule = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ai.metathings.service.deviced.OpModule, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.deviced.OpModule.displayName = 'proto.ai.metathings.service.deviced.OpModule';
}


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
proto.ai.metathings.service.deviced.OpModule.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.deviced.OpModule.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.deviced.OpModule} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.OpModule.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = msg.getId()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    state: jspb.Message.getFieldWithDefault(msg, 2, 0),
    deviceId: (f = msg.getDeviceId()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    endpoint: (f = msg.getEndpoint()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    component: (f = msg.getComponent()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    name: (f = msg.getName()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    alias: (f = msg.getAlias()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    heartbeatAt: (f = msg.getHeartbeatAt()) && proto.google.protobuf.Timestamp.toObject(includeInstance, f)
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
 * @return {!proto.ai.metathings.service.deviced.OpModule}
 */
proto.ai.metathings.service.deviced.OpModule.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.deviced.OpModule;
  return proto.ai.metathings.service.deviced.OpModule.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.deviced.OpModule} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.deviced.OpModule}
 */
proto.ai.metathings.service.deviced.OpModule.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {!proto.ai.metathings.constant.state.ModuleState} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 3:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setDeviceId(value);
      break;
    case 4:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setEndpoint(value);
      break;
    case 5:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setComponent(value);
      break;
    case 6:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setName(value);
      break;
    case 7:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setAlias(value);
      break;
    case 8:
      var value = new proto.google.protobuf.Timestamp;
      reader.readMessage(value,proto.google.protobuf.Timestamp.deserializeBinaryFromReader);
      msg.setHeartbeatAt(value);
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
proto.ai.metathings.service.deviced.OpModule.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.deviced.OpModule.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.deviced.OpModule} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.OpModule.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getDeviceId();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getEndpoint();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getComponent();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getName();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getAlias();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getHeartbeatAt();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      proto.google.protobuf.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.StringValue id = 1;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getId = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 1));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setId = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearId = function() {
  this.setId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasId = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional ai.metathings.constant.state.ModuleState state = 2;
 * @return {!proto.ai.metathings.constant.state.ModuleState}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getState = function() {
  return /** @type {!proto.ai.metathings.constant.state.ModuleState} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.ai.metathings.constant.state.ModuleState} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setState = function(value) {
  jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional google.protobuf.StringValue device_id = 3;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getDeviceId = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 3));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setDeviceId = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearDeviceId = function() {
  this.setDeviceId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasDeviceId = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.StringValue endpoint = 4;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getEndpoint = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 4));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setEndpoint = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearEndpoint = function() {
  this.setEndpoint(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasEndpoint = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.StringValue component = 5;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getComponent = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 5));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setComponent = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearComponent = function() {
  this.setComponent(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasComponent = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional google.protobuf.StringValue name = 6;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getName = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 6));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setName = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearName = function() {
  this.setName(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasName = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional google.protobuf.StringValue alias = 7;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getAlias = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 7));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setAlias = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearAlias = function() {
  this.setAlias(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasAlias = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional google.protobuf.Timestamp heartbeat_at = 8;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.getHeartbeatAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.Timestamp, 8));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.ai.metathings.service.deviced.OpModule.prototype.setHeartbeatAt = function(value) {
  jspb.Message.setWrapperField(this, 8, value);
};


proto.ai.metathings.service.deviced.OpModule.prototype.clearHeartbeatAt = function() {
  this.setHeartbeatAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpModule.prototype.hasHeartbeatAt = function() {
  return jspb.Message.getField(this, 8) != null;
};

