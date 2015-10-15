
export default class ArdaActionCreator {
  constructor(ardaComponent) {
    this.dispatch = ardaComponent.dispatch.bind(ardaComponent)
  }
}

