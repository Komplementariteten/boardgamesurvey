import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ApiSession, Session } from './session';

@Injectable({
  providedIn: 'root'
})
export class SessionService {

  private SessionCookieName: string = "SUS_";

  constructor(private httpClient: HttpClient) { }

  loadInitial()  {
    return this.httpClient.get<ApiSession>("/survey/7jdksjf8238jfwajwf343");
  }
  save(s: Session) {
    const js = s.ToApiSession();
    return;
  }
  createNew() {
    const s =  new Session();
    return s;
  }
  public getSessionId():string {
    let c: string;
    let ca: Array<string> = document.cookie.split(";");
    for (let index = 0; index < ca.length; index++) {
      const element = ca[index];
      c = element.replace("^/\s+/g", '');
      if(c.indexOf(this.SessionCookieName) == 0) {
        console.log(c);
        return c.substring(this.SessionCookieName.length + 1, c.length);
      }
    }
    return "";
  }
}
