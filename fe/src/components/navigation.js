import { useState } from 'react';

import Panel from '../ui/panel';

const Navigation = () => {
    const [showFiles, setShowFiles] = useState(false);

    return (
        <div className="d-flex flex-column bg-dark text-white py-3" >
            <button className="btn btn-dark" onClick={() => setShowFiles(!showFiles)}>
                <i className="fa-regular fa-folder" />
            </button>
            {
                showFiles &&
                    <Panel title="Files" className="mx-2 my-2">
                        <ul className="list-unstyled">
                            <li>File 1</li>
                            <li>File 2</li>
                            <li>File 3</li>
                        </ul>
                    </Panel>
            }
            <span className="spacer flex-grow-1" />
            <button className="btn btn-dark">
                <i className="fa-solid fa-circle-info" />
            </button>
        </div>
    )
}

export default Navigation;