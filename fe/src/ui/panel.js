import { useState, useEffect } from 'react';

const Panel = ({ children, title, onReload, className }) => {
    const [reloading, setReloading] = useState(false);

    useEffect(() => {
        if (reloading && onReload) {
            onReload();
            setTimeout(() => {
                setReloading(false);
            }, 2000);
        }
    }, [reloading, onReload]);

    return (
        <div className={`card ${className}`}>
            <div className="card-header d-flex align-items-center justify-content-between">
                {title}
                { onReload &&
                    <button className={`btn btn-sm btn-outline-primary ${reloading ? 'disabled' : ''}`} onClick={() => setReloading(true)}>
                        { !reloading ?
                            <i className="bi bi-arrow-clockwise"></i> :
                            <div className="spinner-border spinner-border-sm" role="status">
                                <span className="visually-hidden">Reloading...</span>
                            </div>
                        }
                    </button>
                }
            </div>
            <div className="card-body">
                {children}
            </div>
        </div>
    );
}

export default Panel;