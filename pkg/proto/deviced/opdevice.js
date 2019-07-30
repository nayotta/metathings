/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.deviced.OpDevice');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ai.metathings.service.deviced.OpFlow');
goog.require('proto.ai.metathings.service.deviced.OpModule');
goog.require('proto.google.protobuf.StringValue');
goog.require('proto.google.protobuf.Timestamp');

goog.forwardDeclare('proto.ai.metathings.constant.kind.DeviceKind');
goog.forwardDeclare('proto.ai.metathings.constant.state.DeviceState');

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
proto.ai.metathings.service.deviced.OpDevice = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ai.metathings.service.deviced.OpDevice.repeatedFields_, null);
};
goog.inherits(proto.ai.metathings.service.deviced.OpDevice, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.deviced.OpDevice.displayName = 'proto.ai.metathings.service.deviced.OpDevice';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ai.metathings.service.deviced.OpDevice.repeatedFields_ = [6,8];



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
proto.ai.metathings.service.deviced.OpDevice.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.deviced.OpDevice.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.deviced.OpDevice} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.OpDevice.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = msg.getId()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    kind: jspb.Message.getFieldWithDefault(msg, 2, 0),
    state: jspb.Message.getFieldWithDefault(msg, 3, 0),
    name: (f = msg.getName()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    alias: (f = msg.getAlias()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    modulesList: jspb.Message.toObjectList(msg.getModulesList(),
    proto.ai.metathings.service.deviced.OpModule.toObject, includeInstance),
    heartbeatAt: (f = msg.getHeartbeatAt()) && proto.google.protobuf.Timestamp.toObject(includeInstance, f),
    flowsList: jspb.Message.toObjectList(msg.getFlowsList(),
    proto.ai.metathings.service.deviced.OpFlow.toObject, includeInstance)
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
 * @return {!proto.ai.metathings.service.deviced.OpDevice}
 */
proto.ai.metathings.service.deviced.OpDevice.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.deviced.OpDevice;
  return proto.ai.metathings.service.deviced.OpDevice.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.deviced.OpDevice} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.deviced.OpDevice}
 */
proto.ai.metathings.service.deviced.OpDevice.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {!proto.ai.metathings.constant.kind.DeviceKind} */ (reader.readEnum());
      msg.setKind(value);
      break;
    case 3:
      var value = /** @type {!proto.ai.metathings.constant.state.DeviceState} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 4:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setName(value);
      break;
    case 5:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setAlias(value);
      break;
    case 6:
      var value = new proto.ai.metathings.service.deviced.OpModule;
      reader.readMessage(value,proto.ai.metathings.service.deviced.OpModule.deserializeBinaryFromReader);
      msg.addModules(value);
      break;
    case 7:
      var value = new proto.google.protobuf.Timestamp;
      reader.readMessage(value,proto.google.protobuf.Timestamp.deserializeBinaryFromReader);
      msg.setHeartbeatAt(value);
      break;
    case 8:
      var value = new proto.ai.metathings.service.deviced.OpFlow;
      reader.readMessage(value,proto.ai.metathings.service.deviced.OpFlow.deserializeBinaryFromReader);
      msg.addFlows(value);
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
proto.ai.metathings.service.deviced.OpDevice.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.deviced.OpDevice.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.deviced.OpDevice} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.OpDevice.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getKind();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getName();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getAlias();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getModulesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      6,
      f,
      proto.ai.metathings.service.deviced.OpModule.serializeBinaryToWriter
    );
  }
  f = message.getHeartbeatAt();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      proto.google.protobuf.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getFlowsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      proto.ai.metathings.service.deviced.OpFlow.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.StringValue id = 1;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getId = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 1));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setId = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.ai.metathings.service.deviced.OpDevice.prototype.clearId = function() {
  this.setId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.hasId = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional ai.metathings.constant.kind.DeviceKind kind = 2;
 * @return {!proto.ai.metathings.constant.kind.DeviceKind}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getKind = function() {
  return /** @type {!proto.ai.metathings.constant.kind.DeviceKind} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.ai.metathings.constant.kind.DeviceKind} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setKind = function(value) {
  jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional ai.metathings.constant.state.DeviceState state = 3;
 * @return {!proto.ai.metathings.constant.state.DeviceState}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getState = function() {
  return /** @type {!proto.ai.metathings.constant.state.DeviceState} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {!proto.ai.metathings.constant.state.DeviceState} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setState = function(value) {
  jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional google.protobuf.StringValue name = 4;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getName = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 4));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setName = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.ai.metathings.service.deviced.OpDevice.prototype.clearName = function() {
  this.setName(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.hasName = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.StringValue alias = 5;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getAlias = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 5));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setAlias = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.ai.metathings.service.deviced.OpDevice.prototype.clearAlias = function() {
  this.setAlias(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.hasAlias = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated OpModule modules = 6;
 * @return {!Array<!proto.ai.metathings.service.deviced.OpModule>}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getModulesList = function() {
  return /** @type{!Array<!proto.ai.metathings.service.deviced.OpModule>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ai.metathings.service.deviced.OpModule, 6));
};


/** @param {!Array<!proto.ai.metathings.service.deviced.OpModule>} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setModulesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 6, value);
};


/**
 * @param {!proto.ai.metathings.service.deviced.OpModule=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ai.metathings.service.deviced.OpModule}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.addModules = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 6, opt_value, proto.ai.metathings.service.deviced.OpModule, opt_index);
};


proto.ai.metathings.service.deviced.OpDevice.prototype.clearModulesList = function() {
  this.setModulesList([]);
};


/**
 * optional google.protobuf.Timestamp heartbeat_at = 7;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getHeartbeatAt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.Timestamp, 7));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setHeartbeatAt = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


proto.ai.metathings.service.deviced.OpDevice.prototype.clearHeartbeatAt = function() {
  this.setHeartbeatAt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.hasHeartbeatAt = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * repeated OpFlow flows = 8;
 * @return {!Array<!proto.ai.metathings.service.deviced.OpFlow>}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.getFlowsList = function() {
  return /** @type{!Array<!proto.ai.metathings.service.deviced.OpFlow>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ai.metathings.service.deviced.OpFlow, 8));
};


/** @param {!Array<!proto.ai.metathings.service.deviced.OpFlow>} value */
proto.ai.metathings.service.deviced.OpDevice.prototype.setFlowsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.ai.metathings.service.deviced.OpFlow=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ai.metathings.service.deviced.OpFlow}
 */
proto.ai.metathings.service.deviced.OpDevice.prototype.addFlows = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.ai.metathings.service.deviced.OpFlow, opt_index);
};


proto.ai.metathings.service.deviced.OpDevice.prototype.clearFlowsList = function() {
  this.setFlowsList([]);
};


