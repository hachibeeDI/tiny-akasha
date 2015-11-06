import Arda from 'arda';

import jwt from 'jsonwebtoken';


let PostFrontComponent = React.createClass({
  mixins: [Arda.mixin, React.addons.LinkedStateMixin],
  getInitialState() {
    return {};
  },

  onClickLogout() {
    this.dispatch('logout');
  },

  render() {
    let jwtPayload = this.props.localToken === ''? {} : jwt.decode(this.props.localToken);
    if (jwtPayload.user_id) {
      return (
        <a
          className='button__login-state--sign-uped'
          href='/'
          data-bg-image={jwtPayload.image_url}
          onClick={this.onClickLogout}>Logout
        </a>
      );
    } else {
      return (
        <a className='button__login-state--sign-up' href='/login'>Login</a>
      );
    }
  }
});


/*
* 投稿用パネルの各項目などを管理する
*/
export default class PostPanelContext extends Arda.Context {
  get component() {
    return PostFrontComponent;
  }

  delegate(subscribe) {
    super.delegate();
    subscribe('logout', () => {
      localStorage.setItem('localToken', '');
    });
  }

  initState(props) {
    return {localToken: localStorage.getItem('localToken') || ''};
  }

  expandComponentProps(props, state) {
    return state;
  }
}


