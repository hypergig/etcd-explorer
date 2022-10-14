import Toolbar from '@mui/material/Toolbar';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import logo from './assets/images/logo-etcd.png';
import EtcdTree from './EtcdTree';

const drawerWidth = 240;

export default function LeftNav() {
    return (
        <Drawer
            sx={{
                width: drawerWidth,
                flexShrink: 0,
                '& .MuiDrawer-paper': {
                    width: drawerWidth,
                    boxSizing: 'border-box',
                },
            }}
            variant="permanent"
            anchor="left"
        >
            <Toolbar>
                <img src={logo} id="logo" alt="logo"/>
            </Toolbar>
            <Divider/>
            <EtcdTree/>
        </Drawer>
    )
}
