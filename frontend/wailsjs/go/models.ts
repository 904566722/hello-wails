export namespace api {
	
	export class BaseRequest {
	    data: string;
	    keyType: number;
	    action: number;
	
	    static createFrom(source: any = {}) {
	        return new BaseRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.keyType = source["keyType"];
	        this.action = source["action"];
	    }
	}
	export class BaseResponse {
	    code: number;
	    message: string;
	    data: any;
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new BaseResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	        this.total = source["total"];
	    }
	}
	export class ListRequest {
	    limit: number;
	
	    static createFrom(source: any = {}) {
	        return new ListRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.limit = source["limit"];
	    }
	}

}

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

