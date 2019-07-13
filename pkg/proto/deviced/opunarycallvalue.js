/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.deviced.OpUnaryCallValue');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.google.protobuf.Any');
goog.require('proto.google.protobuf.StringValue');


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
proto.ai.metathings.service.deviced.OpUnaryCallValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ai.metathings.service.deviced.OpUnaryCallValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.deviced.OpUnaryCallValue.displayName = 'proto.ai.metathings.service.deviced.OpUnaryCallValue';
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
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.deviced.OpUnaryCallValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.deviced.OpUnaryCallValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: (f = msg.getName()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    component: (f = msg.getComponent()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    method: (f = msg.getMethod()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    value: (f = msg.getValue()) && proto.google.protobuf.Any.toObject(includeInstance, f)
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
 * @return {!proto.ai.metathings.service.deviced.OpUnaryCallValue}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.deviced.OpUnaryCallValue;
  return proto.ai.metathings.service.deviced.OpUnaryCallValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.deviced.OpUnaryCallValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.deviced.OpUnaryCallValue}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setName(value);
      break;
    case 2:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setComponent(value);
      break;
    case 3:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setMethod(value);
      break;
    case 4:
      var value = new proto.google.protobuf.Any;
      reader.readMessage(value,proto.google.protobuf.Any.deserializeBinaryFromReader);
      msg.setValue(value);
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
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.deviced.OpUnaryCallValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.deviced.OpUnaryCallValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getComponent();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getMethod();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getValue();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.google.protobuf.Any.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.StringValue name = 1;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.getName = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 1));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.setName = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.clearName = function() {
  this.setName(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.hasName = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.StringValue component = 2;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.getComponent = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 2));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.setComponent = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.clearComponent = function() {
  this.setComponent(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.hasComponent = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.StringValue method = 3;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.getMethod = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 3));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.setMethod = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.clearMethod = function() {
  this.setMethod(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.hasMethod = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Any value = 4;
 * @return {?proto.google.protobuf.Any}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.getValue = function() {
  return /** @type{?proto.google.protobuf.Any} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.Any, 4));
};


/** @param {?proto.google.protobuf.Any|undefined} value */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.setValue = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.clearValue = function() {
  this.setValue(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.OpUnaryCallValue.prototype.hasValue = function() {
  return jspb.Message.getField(this, 4) != null;
};


