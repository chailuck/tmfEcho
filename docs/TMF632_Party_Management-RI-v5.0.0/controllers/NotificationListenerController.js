/**
 * The NotificationListenerController file is a very simple one, which does not need to be changed manually,
 * unless there's a case where business logic routes the request to an entity which is not
 * the service.
 * The heavy lifting of the Controller item is done in Request.js - that is where request
 * parameters are extracted and sent to the service, and where response is handled.
 */

const Controller = require('./Controller');
const service = require('../services/NotificationListenerService');
const individualAttributeValueChangeEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.individualAttributeValueChangeEvent);
};

const individualCreateEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.individualCreateEvent);
};

const individualDeleteEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.individualDeleteEvent);
};

const individualStateChangeEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.individualStateChangeEvent);
};

const organizationAttributeValueChangeEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.organizationAttributeValueChangeEvent);
};

const organizationCreateEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.organizationCreateEvent);
};

const organizationDeleteEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.organizationDeleteEvent);
};

const organizationStateChangeEvent = async (request, response) => {
  await Controller.handleRequest(request, response, service.organizationStateChangeEvent);
};


module.exports = {
  individualAttributeValueChangeEvent,
  individualCreateEvent,
  individualDeleteEvent,
  individualStateChangeEvent,
  organizationAttributeValueChangeEvent,
  organizationCreateEvent,
  organizationDeleteEvent,
  organizationStateChangeEvent,
};
