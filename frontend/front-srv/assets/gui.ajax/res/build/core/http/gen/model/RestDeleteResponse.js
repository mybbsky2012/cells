/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

/**
* The RestDeleteResponse model module.
* @module model/RestDeleteResponse
* @version 1.0
*/

var RestDeleteResponse = (function () {
    /**
    * Constructs a new <code>RestDeleteResponse</code>.
    * @alias module:model/RestDeleteResponse
    * @class
    */

    function RestDeleteResponse() {
        _classCallCheck(this, RestDeleteResponse);

        this.Success = undefined;
        this.NumRows = undefined;
    }

    /**
    * Constructs a <code>RestDeleteResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestDeleteResponse} obj Optional instance to populate.
    * @return {module:model/RestDeleteResponse} The populated <code>RestDeleteResponse</code> instance.
    */

    RestDeleteResponse.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestDeleteResponse();

            if (data.hasOwnProperty('Success')) {
                obj['Success'] = _ApiClient2['default'].convertToType(data['Success'], 'Boolean');
            }
            if (data.hasOwnProperty('NumRows')) {
                obj['NumRows'] = _ApiClient2['default'].convertToType(data['NumRows'], 'String');
            }
        }
        return obj;
    };

    /**
    * @member {Boolean} Success
    */
    return RestDeleteResponse;
})();

exports['default'] = RestDeleteResponse;
module.exports = exports['default'];

/**
* @member {String} NumRows
*/