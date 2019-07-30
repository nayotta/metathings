/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.policyd.NewAdapterRequest');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');


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
proto.ai.metathings.service.policyd.NewAdapterRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ai.metathings.service.policyd.NewAdapterRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.policyd.NewAdapterRequest.displayName = 'proto.ai.metathings.service.policyd.NewAdapterRequest';
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
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.policyd.NewAdapterRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.policyd.NewAdapterRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    adaptername: jspb.Message.getFieldWithDefault(msg, 1, ""),
    drivername: jspb.Message.getFieldWithDefault(msg, 2, ""),
    connectstring: jspb.Message.getFieldWithDefault(msg, 3, "")
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
 * @return {!proto.ai.metathings.service.policyd.NewAdapterRequest}
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.policyd.NewAdapterRequest;
  return proto.ai.metathings.service.policyd.NewAdapterRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.policyd.NewAdapterRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.policyd.NewAdapterRequest}
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAdaptername(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setDrivername(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setConnectstring(value);
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
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.policyd.NewAdapterRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.policyd.NewAdapterRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAdaptername();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDrivername();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getConnectstring();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string adapterName = 1;
 * @return {string}
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.getAdaptername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.setAdaptername = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string driverName = 2;
 * @return {string}
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.getDrivername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.setDrivername = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string connectString = 3;
 * @return {string}
 */
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.getConnectstring = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.ai.metathings.service.policyd.NewAdapterRequest.prototype.setConnectstring = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


