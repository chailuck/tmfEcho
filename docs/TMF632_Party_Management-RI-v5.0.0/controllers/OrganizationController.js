/**
 * The OrganizationController file is a very simple one, which does not need to be changed manually,
 * unless there's a case where business logic routes the request to an entity which is not
 * the service.
 * The heavy lifting of the Controller item is done in Request.js - that is where request
 * parameters are extracted and sent to the service, and where response is handled.
 */

const Controller = require('./Controller');
const service = require('../services/OrganizationService');
const createOrganization = async (request, response) => {
  await Controller.handleRequest(request, response, service.createOrganization);
};

const deleteOrganization = async (request, response) => {
  await Controller.handleRequest(request, response, service.deleteOrganization);
};

const listOrganization = async (request, response) => {
  await Controller.handleRequest(request, response, service.listOrganization);
};

const patchOrganization = async (request, response) => {
  await Controller.handleRequest(request, response, service.patchOrganization);
};

const retrieveOrganization = async (request, response) => {
  await Controller.handleRequest(request, response, service.retrieveOrganization);
};


module.exports = {
  createOrganization,
  deleteOrganization,
  listOrganization,
  patchOrganization,
  retrieveOrganization,
};
