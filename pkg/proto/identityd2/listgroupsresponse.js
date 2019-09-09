/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.identityd2.ListGroupsResponse');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ai.metathings.service.identityd2.Group');


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
proto.ai.metathings.service.identityd2.ListGroupsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ai.metathings.service.identityd2.ListGroupsResponse.repeatedFields_, null);
};
goog.inherits(proto.ai.metathings.service.identityd2.ListGroupsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.identityd2.ListGroupsResponse.displayName = 'proto.ai.metathings.service.identityd2.ListGroupsResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.repeatedFields_ = [1];



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
proto.ai.metathings.service.identityd2.ListGroupsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.identityd2.ListGroupsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.identityd2.ListGroupsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    groupsList: jspb.Message.toObjectList(msg.getGroupsList(),
    proto.ai.metathings.service.identityd2.Group.toObject, includeInstance)
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
 * @return {!proto.ai.metathings.service.identityd2.ListGroupsResponse}
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.identityd2.ListGroupsResponse;
  return proto.ai.metathings.service.identityd2.ListGroupsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.identityd2.ListGroupsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.identityd2.ListGroupsResponse}
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ai.metathings.service.identityd2.Group;
      reader.readMessage(value,proto.ai.metathings.service.identityd2.Group.deserializeBinaryFromReader);
      msg.addGroups(value);
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
proto.ai.metathings.service.identityd2.ListGroupsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.identityd2.ListGroupsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.identityd2.ListGroupsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGroupsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.ai.metathings.service.identityd2.Group.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Group groups = 1;
 * @return {!Array<!proto.ai.metathings.service.identityd2.Group>}
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.prototype.getGroupsList = function() {
  return /** @type{!Array<!proto.ai.metathings.service.identityd2.Group>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ai.metathings.service.identityd2.Group, 1));
};


/** @param {!Array<!proto.ai.metathings.service.identityd2.Group>} value */
proto.ai.metathings.service.identityd2.ListGroupsResponse.prototype.setGroupsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.ai.metathings.service.identityd2.Group=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ai.metathings.service.identityd2.Group}
 */
proto.ai.metathings.service.identityd2.ListGroupsResponse.prototype.addGroups = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.ai.metathings.service.identityd2.Group, opt_index);
};


proto.ai.metathings.service.identityd2.ListGroupsResponse.prototype.clearGroupsList = function() {
  this.setGroupsList([]);
};

