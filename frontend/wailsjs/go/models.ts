export namespace main {
	
	export class Web {
	    website: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Web(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.website = source["website"];
	        this.name = source["name"];
	    }
	}

}

