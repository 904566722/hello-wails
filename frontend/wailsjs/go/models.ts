export namespace main {
	
	export class Person {
	    name: string;
	    nickName: string;
	
	    static createFrom(source: any = {}) {
	        return new Person(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.nickName = source["nickName"];
	    }
	}

}

export namespace models {
	
	export class KeyVal {
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyVal(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}

}

