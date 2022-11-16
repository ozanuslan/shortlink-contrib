// source: domain/sitemap/v1/sitemap.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

var tagger_tagger_pb = require('../../../tagger/tagger_pb.js');
goog.object.extend(proto, tagger_tagger_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.domain.sitemap.v1.Sitemap', null, global);
goog.exportSymbol('proto.domain.sitemap.v1.Url', null, global);
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
proto.domain.sitemap.v1.Url = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.domain.sitemap.v1.Url, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.domain.sitemap.v1.Url.displayName = 'proto.domain.sitemap.v1.Url';
}
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
proto.domain.sitemap.v1.Sitemap = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.domain.sitemap.v1.Sitemap.repeatedFields_, null);
};
goog.inherits(proto.domain.sitemap.v1.Sitemap, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.domain.sitemap.v1.Sitemap.displayName = 'proto.domain.sitemap.v1.Sitemap';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.domain.sitemap.v1.Url.prototype.toObject = function(opt_includeInstance) {
  return proto.domain.sitemap.v1.Url.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.domain.sitemap.v1.Url} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.domain.sitemap.v1.Url.toObject = function(includeInstance, msg) {
  var f, obj = {
    loc: jspb.Message.getFieldWithDefault(msg, 1, ""),
    lastMod: jspb.Message.getFieldWithDefault(msg, 2, ""),
    changeFreq: jspb.Message.getFieldWithDefault(msg, 3, ""),
    priority: jspb.Message.getFloatingPointFieldWithDefault(msg, 4, 0.0)
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
 * @return {!proto.domain.sitemap.v1.Url}
 */
proto.domain.sitemap.v1.Url.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.domain.sitemap.v1.Url;
  return proto.domain.sitemap.v1.Url.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.domain.sitemap.v1.Url} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.domain.sitemap.v1.Url}
 */
proto.domain.sitemap.v1.Url.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setLoc(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastMod(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setChangeFreq(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setPriority(value);
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
proto.domain.sitemap.v1.Url.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.domain.sitemap.v1.Url.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.domain.sitemap.v1.Url} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.domain.sitemap.v1.Url.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLoc();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getLastMod();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getChangeFreq();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getPriority();
  if (f !== 0.0) {
    writer.writeFloat(
      4,
      f
    );
  }
};


/**
 * optional string loc = 1;
 * @return {string}
 */
proto.domain.sitemap.v1.Url.prototype.getLoc = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.domain.sitemap.v1.Url} returns this
 */
proto.domain.sitemap.v1.Url.prototype.setLoc = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string last_mod = 2;
 * @return {string}
 */
proto.domain.sitemap.v1.Url.prototype.getLastMod = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.domain.sitemap.v1.Url} returns this
 */
proto.domain.sitemap.v1.Url.prototype.setLastMod = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string change_freq = 3;
 * @return {string}
 */
proto.domain.sitemap.v1.Url.prototype.getChangeFreq = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.domain.sitemap.v1.Url} returns this
 */
proto.domain.sitemap.v1.Url.prototype.setChangeFreq = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional float priority = 4;
 * @return {number}
 */
proto.domain.sitemap.v1.Url.prototype.getPriority = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 4, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.domain.sitemap.v1.Url} returns this
 */
proto.domain.sitemap.v1.Url.prototype.setPriority = function(value) {
  return jspb.Message.setProto3FloatField(this, 4, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.domain.sitemap.v1.Sitemap.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.domain.sitemap.v1.Sitemap.prototype.toObject = function(opt_includeInstance) {
  return proto.domain.sitemap.v1.Sitemap.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.domain.sitemap.v1.Sitemap} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.domain.sitemap.v1.Sitemap.toObject = function(includeInstance, msg) {
  var f, obj = {
    urlList: jspb.Message.toObjectList(msg.getUrlList(),
    proto.domain.sitemap.v1.Url.toObject, includeInstance)
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
 * @return {!proto.domain.sitemap.v1.Sitemap}
 */
proto.domain.sitemap.v1.Sitemap.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.domain.sitemap.v1.Sitemap;
  return proto.domain.sitemap.v1.Sitemap.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.domain.sitemap.v1.Sitemap} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.domain.sitemap.v1.Sitemap}
 */
proto.domain.sitemap.v1.Sitemap.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.domain.sitemap.v1.Url;
      reader.readMessage(value,proto.domain.sitemap.v1.Url.deserializeBinaryFromReader);
      msg.addUrl(value);
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
proto.domain.sitemap.v1.Sitemap.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.domain.sitemap.v1.Sitemap.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.domain.sitemap.v1.Sitemap} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.domain.sitemap.v1.Sitemap.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUrlList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.domain.sitemap.v1.Url.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Url url = 1;
 * @return {!Array<!proto.domain.sitemap.v1.Url>}
 */
proto.domain.sitemap.v1.Sitemap.prototype.getUrlList = function() {
  return /** @type{!Array<!proto.domain.sitemap.v1.Url>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.domain.sitemap.v1.Url, 1));
};


/**
 * @param {!Array<!proto.domain.sitemap.v1.Url>} value
 * @return {!proto.domain.sitemap.v1.Sitemap} returns this
*/
proto.domain.sitemap.v1.Sitemap.prototype.setUrlList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.domain.sitemap.v1.Url=} opt_value
 * @param {number=} opt_index
 * @return {!proto.domain.sitemap.v1.Url}
 */
proto.domain.sitemap.v1.Sitemap.prototype.addUrl = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.domain.sitemap.v1.Url, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.domain.sitemap.v1.Sitemap} returns this
 */
proto.domain.sitemap.v1.Sitemap.prototype.clearUrlList = function() {
  return this.setUrlList([]);
};


goog.object.extend(exports, proto.domain.sitemap.v1);
