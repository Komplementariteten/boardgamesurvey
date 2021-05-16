import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { tap } from 'rxjs/operators'
import { SearchItem } from './searchitem';
import { BoardGame } from './boardgame';

@Injectable({
  providedIn: 'root'
})
export class BbgService {
  bbgapiSearchUrl: string = "http://localhost:6100/bbg/search"
  bbgapiDetailsUrl: string = "http://localhost:6100/bbg/game/"
  constructor(private httpClient: HttpClient) {   }
  search(name: string) {
    const sUrl = this.bbgapiSearchUrl + "?name=" + name;
    return this.httpClient.get<SearchItem[]>(sUrl);
  }
  getDetails(id:number, search:string) {
    const sUrl = this.bbgapiDetailsUrl + id;
    return this.httpClient.get<BoardGame>(sUrl).pipe(tap(g => {
      g.Names?.forEach(s => {
        if(s.toLowerCase().indexOf(search) >= 0){
          g.Name = s;
        }
      });
    }));
  }
}
