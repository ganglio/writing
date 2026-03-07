import { useState } from 'react';

import GlobalContext from './context/global';

import Layout from './ui/layout';

import './App.css';

function App() {

    const [random, setRandom] = useState("");

    return (
        <GlobalContext.Provider>
            <Layout />
        </GlobalContext.Provider>
    );
}

export default App;
