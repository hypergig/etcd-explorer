import * as React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import TreeView from '@mui/lab/TreeView';
import TreeItem from '@mui/lab/TreeItem';
import {EventsOn, EventsOnce, LogError, LogInfo} from "../wailsjs/runtime/runtime";
import {StartService} from "../wailsjs/go/listwatcher/Service";
import {etcdtree} from "../wailsjs/go/models";

const rootToTree = (root: { [key: string]: etcdtree.Node }): JSX.Element[] => {
    return Object.values(root).map((node) => toTree(node))
}

const toTree = (node: etcdtree.Node): JSX.Element => {
    if (node.subTree === undefined) {
        return (
            <TreeItem nodeId={node.path} key={node.path} label={node.name}/>
        )
    }
    return (
        <TreeItem nodeId={node.path} key={node.path} label={node.name}>
            {rootToTree(node.subTree)}
        </TreeItem>
    )
};

const start = () => {
    StartService("localhost:2379").then((err) => {
        if (err !== undefined) {
            LogError(err.message)
        } else {
            LogInfo("service started")
        }
    });
}

export default function EtcdTree() {
    const [root, setRoot] = React.useState<{ [key: string]: etcdtree.Node }>({});
    React.useEffect(() => {
        EventsOn("etcEvent", (payload: { [key: string]: etcdtree.Node }) => setRoot(payload));
        LogInfo("registered handler");
    }, [])
    return (
        <Box sx={{height: 270, flexGrow: 1, maxWidth: 400, overflowY: 'auto'}}>
            <Box sx={{mb: 1}}>
                <Button variant="contained" color="success" onClick={start}>
                    Start
                </Button>
            </Box>

            <TreeView
                aria-label="file system navigator"
                defaultCollapseIcon={<ExpandMoreIcon/>}
                defaultExpandIcon={<ChevronRightIcon/>}
            >
                {rootToTree(root)}
            </TreeView>
        </Box>
    );
}
