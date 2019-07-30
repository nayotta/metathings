/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.identityd2.OpAction');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Map');
goog.require('jspb.Message');
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
proto.ai.metathings.service.identityd2.OpAction = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ai.metathings.service.identityd2.OpAction, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.identityd2.OpAction.displayName = 'proto.ai.metathings.service.identityd2.OpAction';
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
proto.ai.metathings.service.identityd2.OpAction.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.identityd2.OpAction.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.identityd2.OpAction} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.identityd2.OpAction.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = msg.getId()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    name: (f = msg.getName()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    alias: (f = msg.getAlias()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    description: (f = msg.getDescription()) && proto.google.protobuf.StringValue.toObject(includeInstance, f),
    extraMap: (f = msg.getExtraMap()) ? f.toObject(includeInstance, proto.google.protobuf.StringValue.toObject) : []
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
 * @return {!proto.ai.metathings.service.identityd2.OpAction}
 */
proto.ai.metathings.service.identityd2.OpAction.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.identityd2.OpAction;
  return proto.ai.metathings.service.identityd2.OpAction.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.identityd2.OpAction} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.identityd2.OpAction}
 */
proto.ai.metathings.service.identityd2.OpAction.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setName(value);
      break;
    case 3:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setAlias(value);
      break;
    case 4:
      var value = new proto.google.protobuf.StringValue;
      reader.readMessage(value,proto.google.protobuf.StringValue.deserializeBinaryFromReader);
      msg.setDescription(value);
      break;
    case 5:
      var value = msg.getExtraMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.google.protobuf.StringValue.deserializeBinaryFromReader, "");
         });
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
proto.ai.metathings.service.identityd2.OpAction.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.identityd2.OpAction.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.identityd2.OpAction} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.identityd2.OpAction.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getName();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getAlias();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getDescription();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.google.protobuf.StringValue.serializeBinaryToWriter
    );
  }
  f = message.getExtraMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.google.protobuf.StringValue.serializeBinaryToWriter);
  }
};


/**
 * optional google.protobuf.StringValue id = 1;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.getId = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 1));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.identityd2.OpAction.prototype.setId = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.ai.metathings.service.identityd2.OpAction.prototype.clearId = function() {
  this.setId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.hasId = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.StringValue name = 2;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.getName = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 2));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.identityd2.OpAction.prototype.setName = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.ai.metathings.service.identityd2.OpAction.prototype.clearName = function() {
  this.setName(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.hasName = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.StringValue alias = 3;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.getAlias = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 3));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.identityd2.OpAction.prototype.setAlias = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.ai.metathings.service.identityd2.OpAction.prototype.clearAlias = function() {
  this.setAlias(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.hasAlias = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.StringValue description = 4;
 * @return {?proto.google.protobuf.StringValue}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.getDescription = function() {
  return /** @type{?proto.google.protobuf.StringValue} */ (
    jspb.Message.getWrapperField(this, proto.google.protobuf.StringValue, 4));
};


/** @param {?proto.google.protobuf.StringValue|undefined} value */
proto.ai.metathings.service.identityd2.OpAction.prototype.setDescription = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.ai.metathings.service.identityd2.OpAction.prototype.clearDescription = function() {
  this.setDescription(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.hasDescription = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * map<string, google.protobuf.StringValue> extra = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.google.protobuf.StringValue>}
 */
proto.ai.metathings.service.identityd2.OpAction.prototype.getExtraMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.google.protobuf.StringValue>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      proto.google.protobuf.StringValue));
};


proto.ai.metathings.service.identityd2.OpAction.prototype.clearExtraMap = function() {
  this.getExtraMap().clear();
};


