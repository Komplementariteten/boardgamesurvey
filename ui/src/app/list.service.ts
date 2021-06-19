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
  private session: Session = new Session();
  public votes: number = 0;
  private initialized: boolean = false;
  private voteStore: string = "bgsurveyvotes";

  constructor(private sessionService: SessionService, private websocketService: WebsocketService, private bbgService: BbgService) {
    this.items = [];
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
    const voteStr = localStorage.getItem(this.voteStore);
    if (voteStr === null){
      localStorage.setItem(this.voteStore, this.votes.toString());
    } else {
      this.votes = <number><any>voteStr;
    }
    console.log(this.votes);
  }
  private localVote() {
    localStorage.setItem(this.voteStore, this.votes.toString());
    this.votes++;
  }
  canVote(): boolean {
    return this.votes < MaxNumberOfVotes;
  }
  vote(item: SurveyItem){
    let vote = new Vote();
    vote.Title = item.Title!;
    vote.Description = item.Description!;
    vote.Owner = this.sessionService.getSessionId();
    if(item.Ref !== undefined){
      vote.ObjectId = item.Ref.ObjectId!;
    }
    this.websocketService.send(vote);
    this.localVote();
  }
  add(item: SurveyItem) {
    let vote = new Vote();
    vote.Title = item.Title!;
    vote.ObjectId = -1;
    vote.Description = item.Description!;
    vote.Owner = this.sessionService.getSessionId();
    this.websocketService.send(vote);
    this.localVote();
  }
  addGame(game: BoardGame) {
    let vote = new Vote();
    vote.Title = game.Name!;
    vote.ObjectId = game.ObjectId!;
    vote.Owner = this.sessionService.getSessionId();
    this.websocketService.send(vote);
    this.localVote();
  }
}
