/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.identityd2.DeleteCredentialRequest');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ai.metathings.service.identityd2.OpCredential');


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
proto.ai.metathings.service.identityd2.DeleteCredentialRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ai.metathings.service.identityd2.DeleteCredentialRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.identityd2.DeleteCredentialRequest.displayName = 'proto.ai.metathings.service.identityd2.DeleteCredentialRequest';
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
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.identityd2.DeleteCredentialRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.identityd2.DeleteCredentialRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    credential: (f = msg.getCredential()) && proto.ai.metathings.service.identityd2.OpCredential.toObject(includeInstance, f)
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
 * @return {!proto.ai.metathings.service.identityd2.DeleteCredentialRequest}
 */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.identityd2.DeleteCredentialRequest;
  return proto.ai.metathings.service.identityd2.DeleteCredentialRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.identityd2.DeleteCredentialRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.identityd2.DeleteCredentialRequest}
 */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ai.metathings.service.identityd2.OpCredential;
      reader.readMessage(value,proto.ai.metathings.service.identityd2.OpCredential.deserializeBinaryFromReader);
      msg.setCredential(value);
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
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.identityd2.DeleteCredentialRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.identityd2.DeleteCredentialRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCredential();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.ai.metathings.service.identityd2.OpCredential.serializeBinaryToWriter
    );
  }
};


/**
 * optional OpCredential credential = 1;
 * @return {?proto.ai.metathings.service.identityd2.OpCredential}
 */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.prototype.getCredential = function() {
  return /** @type{?proto.ai.metathings.service.identityd2.OpCredential} */ (
    jspb.Message.getWrapperField(this, proto.ai.metathings.service.identityd2.OpCredential, 1));
};


/** @param {?proto.ai.metathings.service.identityd2.OpCredential|undefined} value */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.prototype.setCredential = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.ai.metathings.service.identityd2.DeleteCredentialRequest.prototype.clearCredential = function() {
  this.setCredential(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.identityd2.DeleteCredentialRequest.prototype.hasCredential = function() {
  return jspb.Message.getField(this, 1) != null;
};


