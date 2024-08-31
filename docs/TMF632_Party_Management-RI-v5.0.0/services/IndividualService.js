/* eslint-disable no-unused-vars */
const Service = require('./Service');

  /**
   * Creates a Individual
   * This operation creates a Individual entity.
   *
   * individualFVO IndividualFVO The Individual to be created
   * fields String Comma-separated properties to be provided in response (optional)
   * returns Individual
   **/
  const createIndividual =  ( args, context /* individualFVOfields  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "Individual";
        context.operationId  = "createIndividual";
        context.method       = "post";
        try {
          /* matching isRestfulCreate */
          resolve(Service.create(args, context ));

        } catch (e) {
          console.log("createIndividual: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Deletes a Individual
   * This operation deletes a Individual entity.
   *
   * id String Identifier of the Resource
   * no response value expected for this operation
   **/
  const deleteIndividual =  ( args, context /* id  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "Individual";
        context.operationId  = "deleteIndividual";
        context.method       = "delete";
        try {
          /* matching isRestfulDestroy */
          resolve(Service.remove(args, context ));

        } catch (e) {
          console.log("deleteIndividual: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * List or find Individual objects
   * List or find Individual objects
   *
   * fields String Comma-separated properties to be provided in response (optional)
   * offset Integer Requested index for start of resources to be provided in response (optional)
   * limit Integer Requested number of resources to be provided in response (optional)
   * returns List
   **/
  const listIndividual =  ( args, context /* fieldsoffsetlimit  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "Individual";
        context.operationId  = "listIndividual";
        context.method       = "get";
        try {
          /* matching isRestfulIndex */
          resolve(Service.index(args, context ));

        } catch (e) {
          console.log("listIndividual: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Updates partially a Individual
   * This operation updates partially a Individual entity.
   *
   * id String Identifier of the Resource
   * individualMVO IndividualMVO The Individual to be patched
   * fields String Comma-separated properties to be provided in response (optional)
   * returns Individual
   **/
  const patchIndividual =  ( args, context /* idindividualMVOfields  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "Individual";
        context.operationId  = "patchIndividual";
        context.method       = "patch";
        try {
          /* matching isRestfulUpdate */
          resolve(Service.update(args, context ));

        } catch (e) {
          console.log("patchIndividual: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Retrieves a Individual by ID
   * This operation retrieves a Individual entity. Attribute selection enabled for all first level attributes.
   *
   * id String Identifier of the Resource
   * fields String Comma-separated properties to be provided in response (optional)
   * returns Individual
   **/
  const retrieveIndividual =  ( args, context /* idfields  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "Individual";
        context.operationId  = "retrieveIndividual";
        context.method       = "get";
        try {
          /* matching isRestfulShow */
          resolve(Service.show(args, context ));

        } catch (e) {
          console.log("retrieveIndividual: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    

module.exports = {
  createIndividual,
  deleteIndividual,
  listIndividual,
  patchIndividual,
  retrieveIndividual,
};
