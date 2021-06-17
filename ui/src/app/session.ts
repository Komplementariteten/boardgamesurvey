/* export class Session {
  public readonly Content: Map<string, any> = new Map<string, any>();
} */

export class Session {
  [index: string]: any;
  FromApiSession(a: ApiSession) {
    for(const key in a.Content){
      this[key] = a.Content[key];
    }
  }
  ToApiSession(): ApiSession {
    const api = new ApiSession();
    for(const key in this) {
      api.Content[key] = this[key];
    }
    return api;
  }
  count(): number {
    let count = 0;
    for(const key in this){
      count++;
    }
    return count;
  }
}

export class ApiSession{
  public readonly Collection: string = "LingenBrettspiel300EuroSession";
  public readonly Type: string = "map";
  public readonly Content: any = {};
}
