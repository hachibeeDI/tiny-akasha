
export default class ArdaActionCreator {
  constructor(ardaComponent) {
    this.dispatch = component.dispatch.bind(component)
  }
}

