/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.policyd.Array2DReply');
goog.provide('proto.ai.metathings.service.policyd.Array2DReply.d');

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
proto.ai.metathings.service.policyd.Array2DReply = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ai.metathings.service.policyd.Array2DReply.repeatedFields_, null);
};
goog.inherits(proto.ai.metathings.service.policyd.Array2DReply, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.policyd.Array2DReply.displayName = 'proto.ai.metathings.service.policyd.Array2DReply';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ai.metathings.service.policyd.Array2DReply.repeatedFields_ = [1];



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
proto.ai.metathings.service.policyd.Array2DReply.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.policyd.Array2DReply.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.policyd.Array2DReply} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.policyd.Array2DReply.toObject = function(includeInstance, msg) {
  var f, obj = {
    d2List: jspb.Message.toObjectList(msg.getD2List(),
    proto.ai.metathings.service.policyd.Array2DReply.d.toObject, includeInstance)
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
 * @return {!proto.ai.metathings.service.policyd.Array2DReply}
 */
proto.ai.metathings.service.policyd.Array2DReply.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.policyd.Array2DReply;
  return proto.ai.metathings.service.policyd.Array2DReply.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.policyd.Array2DReply} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.policyd.Array2DReply}
 */
proto.ai.metathings.service.policyd.Array2DReply.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ai.metathings.service.policyd.Array2DReply.d;
      reader.readMessage(value,proto.ai.metathings.service.policyd.Array2DReply.d.deserializeBinaryFromReader);
      msg.addD2(value);
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
proto.ai.metathings.service.policyd.Array2DReply.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.policyd.Array2DReply.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.policyd.Array2DReply} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.policyd.Array2DReply.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getD2List();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.ai.metathings.service.policyd.Array2DReply.d.serializeBinaryToWriter
    );
  }
};



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
proto.ai.metathings.service.policyd.Array2DReply.d = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ai.metathings.service.policyd.Array2DReply.d.repeatedFields_, null);
};
goog.inherits(proto.ai.metathings.service.policyd.Array2DReply.d, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.policyd.Array2DReply.d.displayName = 'proto.ai.metathings.service.policyd.Array2DReply.d';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ai.metathings.service.policyd.Array2DReply.d.repeatedFields_ = [1];



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
proto.ai.metathings.service.policyd.Array2DReply.d.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.policyd.Array2DReply.d.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.policyd.Array2DReply.d} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.policyd.Array2DReply.d.toObject = function(includeInstance, msg) {
  var f, obj = {
    d1List: jspb.Message.getRepeatedField(msg, 1)
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
 * @return {!proto.ai.metathings.service.policyd.Array2DReply.d}
 */
proto.ai.metathings.service.policyd.Array2DReply.d.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.policyd.Array2DReply.d;
  return proto.ai.metathings.service.policyd.Array2DReply.d.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.policyd.Array2DReply.d} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.policyd.Array2DReply.d}
 */
proto.ai.metathings.service.policyd.Array2DReply.d.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addD1(value);
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
proto.ai.metathings.service.policyd.Array2DReply.d.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.policyd.Array2DReply.d.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.policyd.Array2DReply.d} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.policyd.Array2DReply.d.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getD1List();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string d1 = 1;
 * @return {!Array<string>}
 */
proto.ai.metathings.service.policyd.Array2DReply.d.prototype.getD1List = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/** @param {!Array<string>} value */
proto.ai.metathings.service.policyd.Array2DReply.d.prototype.setD1List = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.ai.metathings.service.policyd.Array2DReply.d.prototype.addD1 = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


proto.ai.metathings.service.policyd.Array2DReply.d.prototype.clearD1List = function() {
  this.setD1List([]);
};


/**
 * repeated d d2 = 1;
 * @return {!Array<!proto.ai.metathings.service.policyd.Array2DReply.d>}
 */
proto.ai.metathings.service.policyd.Array2DReply.prototype.getD2List = function() {
  return /** @type{!Array<!proto.ai.metathings.service.policyd.Array2DReply.d>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ai.metathings.service.policyd.Array2DReply.d, 1));
};


/** @param {!Array<!proto.ai.metathings.service.policyd.Array2DReply.d>} value */
proto.ai.metathings.service.policyd.Array2DReply.prototype.setD2List = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.ai.metathings.service.policyd.Array2DReply.d=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ai.metathings.service.policyd.Array2DReply.d}
 */
proto.ai.metathings.service.policyd.Array2DReply.prototype.addD2 = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.ai.metathings.service.policyd.Array2DReply.d, opt_index);
};


proto.ai.metathings.service.policyd.Array2DReply.prototype.clearD2List = function() {
  this.setD2List([]);
};

