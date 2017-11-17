import { Injectable, Inject } from "@angular/core";
import { HttpClient } from "@angular/common/http/src/client";

@Injectable()
export class BackendService {
    protected urlAPI: any;

    constructor( @Inject(HttpClient) protected http?) {
    }

    public gets(limit?, offset?) {

    }
}