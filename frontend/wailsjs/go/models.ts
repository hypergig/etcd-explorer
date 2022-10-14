export namespace etcdtree {
	
	export class Node {
	    name: string;
	    path: string;
	    isKey?: boolean;
	    subTree?: {[key: string]: Node};
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.isKey = source["isKey"];
	        this.subTree = source["subTree"];
	    }
	}
	export class Root {
	    tree: {[key: string]: Node};
	
	    static createFrom(source: any = {}) {
	        return new Root(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tree = source["tree"];
	    }
	}

}

