import { Injectable } from '@angular/core';
import { BbgService } from './bbg.service';
import { BoardGame } from './boardgame';
import { MaxNumberOfVotes, VoteObjectName } from './constants';
import { ApiSession, Session } from './session';
import { SessionService } from './session.service';
import { SurveyItem, Vote, VoteEvent, Votes } from './surveyitem';
import { WebsocketService } from './websocket.service';

@Injectable({
  providedIn: 'root'
})
export class ListService {

  public items: SurveyItem[];
  private games: BoardGame[];
  private session: Session = new Session();
  private votes: string[] = [];
  private initialized: boolean = false;

  constructor(private sessionService: SessionService, private websocketService: WebsocketService, private bbgService: BbgService) {
    this.items = [];
    this.games = [];
    this.websocketService.connect();
    let eventEmiter = this.websocketService.getEventListener();
    eventEmiter.subscribe((ev: VoteEvent) => {
      if (ev.data!.Votes == undefined) {
        return;
      }
      const votes = ev.data?.Votes;
      this.items = [];
      for (let index = 0; index < votes!.length; index++) {
        const si = new SurveyItem ();
        si.Title = votes![index].Title;
        si.Votes = votes![index].Votes;
        if(votes![index].ObjectId >= 0) {
          this.bbgService.getDetails(votes![index].ObjectId, si.Title!).subscribe(d => {
            console.log("Game searched");
            console.log(d);
            si.Ref = d;
          });
        }
        this.items[index] = si;
      }
    });
  }
  initialize(s: ApiSession) {
    if(this.initialized) {
      return;
    }
    this.initialized = true;
    this.session.FromApiSession(s);
    if (this.session[VoteObjectName]) {
      const votes = this.session[VoteObjectName];
      const votesMap = <string[]>votes;
      if (votes != undefined) {
        this.votes = votesMap;
      }
    }
  }
  canVote(): boolean {
    return this.votes.length < MaxNumberOfVotes;
  }
  vote(item: SurveyItem){
    let vote = new Vote();
    vote.Title = item.Title!;
    vote.Owner = this.sessionService.getSessionId();
    if(item.Ref !== undefined){
      vote.ObjectId = item.Ref.ObjectId!;
    }
    this.websocketService.send(vote);
  }
  add(item: SurveyItem) {
    this.items.push(item);
  }
  addGame(game: BoardGame) {
    this.games.push(game);
    let item = new SurveyItem();
    item.Ref = game;
    item.Description = game.Description;
    item.Title = game.Name;
    this.items.push(item);
  }
}
