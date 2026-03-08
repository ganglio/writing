import { useContext } from 'react';

import GlobalContext from '../context/global';

const Modal = ({ title, children, onSave }) => {
    
    const { dispatch } = useContext(GlobalContext.Context);

    const onClose = () => {
        dispatch({ type: GlobalContext.Actions.TOGGLE_MODAL });
    }
    
    return (
        <div className="modal show d-block" tabIndex="-1" aria-hidden="false">
            <div className="modal-dialog modal-fullscreen">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title">{title}</h5>
                        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div className="modal-body">
                        { children }
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary" data-bs-dismiss="modal" onClick={onClose}>Close</button>
                        <button type="button" className="btn btn-primary" onClick={onSave}>Save changes</button>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Modal;