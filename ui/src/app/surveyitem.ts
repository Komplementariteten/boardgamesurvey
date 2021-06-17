import { BoardGame } from './boardgame';

export class SurveyItem {
  public Ref: BoardGame | undefined;
  public Votes: number = 0;
  public Description: string | undefined;
  public Title: string | undefined;
}

export class VoteEvent {
  public type: string = "";
  public data : Votes | undefined;
}

export class Votes {
  public Votes: VoteItem[] | undefined;
}

export class VoteItem {
  public ObjectId: number = 0;
  public Title: string | undefined;
  public Votes: number = 0;
}

export class Vote {
  public Title: string = "";
  public Owner: string = "";
  public ObjectId: number = -1;
}
