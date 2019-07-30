/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse');
goog.provide('proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ai.metathings.service.deviced.Flow');
goog.require('proto.ai.metathings.service.deviced.Frame');


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
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.repeatedFields_, null);
};
goog.inherits(proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.displayName = 'proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.repeatedFields_ = [1];



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
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    packsList: jspb.Message.toObjectList(msg.getPacksList(),
    proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.toObject, includeInstance)
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
 * @return {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse;
  return proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack;
      reader.readMessage(value,proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.deserializeBinaryFromReader);
      msg.addPacks(value);
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
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPacksList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.serializeBinaryToWriter
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
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.repeatedFields_, null);
};
goog.inherits(proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.displayName = 'proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.repeatedFields_ = [2];



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
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.toObject = function(opt_includeInstance) {
  return proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.toObject = function(includeInstance, msg) {
  var f, obj = {
    flow: (f = msg.getFlow()) && proto.ai.metathings.service.deviced.Flow.toObject(includeInstance, f),
    framesList: jspb.Message.toObjectList(msg.getFramesList(),
    proto.ai.metathings.service.deviced.Frame.toObject, includeInstance)
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
 * @return {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack;
  return proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ai.metathings.service.deviced.Flow;
      reader.readMessage(value,proto.ai.metathings.service.deviced.Flow.deserializeBinaryFromReader);
      msg.setFlow(value);
      break;
    case 2:
      var value = new proto.ai.metathings.service.deviced.Frame;
      reader.readMessage(value,proto.ai.metathings.service.deviced.Frame.deserializeBinaryFromReader);
      msg.addFrames(value);
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
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFlow();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.ai.metathings.service.deviced.Flow.serializeBinaryToWriter
    );
  }
  f = message.getFramesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.ai.metathings.service.deviced.Frame.serializeBinaryToWriter
    );
  }
};


/**
 * optional Flow flow = 1;
 * @return {?proto.ai.metathings.service.deviced.Flow}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.getFlow = function() {
  return /** @type{?proto.ai.metathings.service.deviced.Flow} */ (
    jspb.Message.getWrapperField(this, proto.ai.metathings.service.deviced.Flow, 1));
};


/** @param {?proto.ai.metathings.service.deviced.Flow|undefined} value */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.setFlow = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.clearFlow = function() {
  this.setFlow(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.hasFlow = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Frame frames = 2;
 * @return {!Array<!proto.ai.metathings.service.deviced.Frame>}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.getFramesList = function() {
  return /** @type{!Array<!proto.ai.metathings.service.deviced.Frame>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ai.metathings.service.deviced.Frame, 2));
};


/** @param {!Array<!proto.ai.metathings.service.deviced.Frame>} value */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.setFramesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.ai.metathings.service.deviced.Frame=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ai.metathings.service.deviced.Frame}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.addFrames = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.ai.metathings.service.deviced.Frame, opt_index);
};


proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack.prototype.clearFramesList = function() {
  this.setFramesList([]);
};


/**
 * repeated Pack packs = 1;
 * @return {!Array<!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack>}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.prototype.getPacksList = function() {
  return /** @type{!Array<!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack, 1));
};


/** @param {!Array<!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack>} value */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.prototype.setPacksList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack}
 */
proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.prototype.addPacks = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.Pack, opt_index);
};


proto.ai.metathings.service.deviced.QueryFramesFromFlowResponse.prototype.clearPacksList = function() {
  this.setPacksList([]);
};


