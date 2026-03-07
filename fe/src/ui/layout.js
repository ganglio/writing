import { useContext } from 'react';

import GlobalContext from '../context/global';

import Mainarea from '../components/mainarea';
import Tray from '../components/tray';
import Navigation from '../components/navigation';

import modals from '../modals';

const Layout = () => {

    const { state } = useContext(GlobalContext.Context);
    
    const Modal = state.modal ? modals[state.modal] : null;

    return (
        <>
            <div className="container-fluid d-flex h-100 w-100 p-0">
                <Navigation />
                <div className="container-fluid d-flex flex-column px-0">
                    <Mainarea />
                    <Tray />
                </div>
                { Modal ? <Modal /> : null }
            </div>
        </>
    )
}

export default Layout;