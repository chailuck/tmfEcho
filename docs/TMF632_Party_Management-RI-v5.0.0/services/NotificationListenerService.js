/* eslint-disable no-unused-vars */
const Service = require('./Service');

  /**
   * Client listener for entity IndividualAttributeValueChangeEvent
   * Example of a client listener for receiving the notification IndividualAttributeValueChangeEvent
   *
   * individualAttributeValueChangeEvent IndividualAttributeValueChangeEvent Individual attributeValueChange Event payload
   * no response value expected for this operation
   **/
  const individualAttributeValueChangeEvent =  ( args, context /* individualAttributeValueChangeEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "individualAttributeValueChangeEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("individualAttributeValueChangeEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity IndividualCreateEvent
   * Example of a client listener for receiving the notification IndividualCreateEvent
   *
   * individualCreateEvent IndividualCreateEvent Individual create Event payload
   * no response value expected for this operation
   **/
  const individualCreateEvent =  ( args, context /* individualCreateEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "individualCreateEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("individualCreateEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity IndividualDeleteEvent
   * Example of a client listener for receiving the notification IndividualDeleteEvent
   *
   * individualDeleteEvent IndividualDeleteEvent Individual delete Event payload
   * no response value expected for this operation
   **/
  const individualDeleteEvent =  ( args, context /* individualDeleteEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "individualDeleteEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("individualDeleteEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity IndividualStateChangeEvent
   * Example of a client listener for receiving the notification IndividualStateChangeEvent
   *
   * individualStateChangeEvent IndividualStateChangeEvent Individual stateChange Event payload
   * no response value expected for this operation
   **/
  const individualStateChangeEvent =  ( args, context /* individualStateChangeEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "individualStateChangeEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("individualStateChangeEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity OrganizationAttributeValueChangeEvent
   * Example of a client listener for receiving the notification OrganizationAttributeValueChangeEvent
   *
   * organizationAttributeValueChangeEvent OrganizationAttributeValueChangeEvent Organization attributeValueChange Event payload
   * no response value expected for this operation
   **/
  const organizationAttributeValueChangeEvent =  ( args, context /* organizationAttributeValueChangeEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "organizationAttributeValueChangeEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("organizationAttributeValueChangeEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity OrganizationCreateEvent
   * Example of a client listener for receiving the notification OrganizationCreateEvent
   *
   * organizationCreateEvent OrganizationCreateEvent Organization create Event payload
   * no response value expected for this operation
   **/
  const organizationCreateEvent =  ( args, context /* organizationCreateEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "organizationCreateEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("organizationCreateEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity OrganizationDeleteEvent
   * Example of a client listener for receiving the notification OrganizationDeleteEvent
   *
   * organizationDeleteEvent OrganizationDeleteEvent Organization delete Event payload
   * no response value expected for this operation
   **/
  const organizationDeleteEvent =  ( args, context /* organizationDeleteEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "organizationDeleteEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("organizationDeleteEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    
  /**
   * Client listener for entity OrganizationStateChangeEvent
   * Example of a client listener for receiving the notification OrganizationStateChangeEvent
   *
   * organizationStateChangeEvent OrganizationStateChangeEvent Organization stateChange Event payload
   * no response value expected for this operation
   **/
  const organizationStateChangeEvent =  ( args, context /* organizationStateChangeEvent  */) =>
    new Promise(
      async (resolve) => {
        context.classname    = "NotificationListener";
        context.operationId  = "organizationStateChangeEvent";
        context.method       = "post";
        try {
          /* NOT matching isRestful */
          resolve(Service.serve(args, context ));

        } catch (e) {
          console.log("organizationStateChangeEvent: error=" + e);
          resolve(Service.rejectResponse(
            e.message || 'Invalid input',
            e.status || 405,
          ));
        }
      },
    )
    

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
