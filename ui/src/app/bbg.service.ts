import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { parse } from 'fast-xml-parser'
import { tap } from 'rxjs/operators'

@Injectable({
  providedIn: 'root'
})
export class BbgService {
  bbgapiSearchUrl: string = "https://www.boardgamegeek.com/xmlapi/search"
  constructor(private httpClient: HttpClient) {   }
  search(name: string) {
    const sUrl = this.bbgapiSearchUrl + "?search=" + name;
    console.log("searching");
    console.log(sUrl);
    return this.httpClient.get(sUrl).pipe(tap(d => {
      console.log(d);
    }));
  }
}
