import Mainarea from '../components/mainarea';
import Tray from '../components/tray';
import Navigation from '../components/navigation';

const Layout = () => {
    return (
        <div className="container-fluid d-flex h-100 w-100 p-0">
            <Navigation />
            <div className="container d-flex flex-column">
                <Mainarea />
                <Tray />
            </div>
        </div>
    )
}

export default Layout;