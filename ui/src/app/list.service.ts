import { Injectable } from '@angular/core';
import { BoardGame } from './boardgame';
import { SurveyItem } from './surveyitem';

@Injectable({
  providedIn: 'root'
})
export class ListService {

  public items: SurveyItem[];
  private games: BoardGame[];
  private hasVoted: boolean = false;
  constructor() {
    this.items = [];
    this.games = [];
  }
  canVote() {
    console.log("can vote");
    return !this.hasVoted;
  }
  vote(item: SurveyItem){
    this.items.forEach(si => {
      if (si == item){
        console.log("voting");
        si.Votes++;
      }
    });
    this.hasVoted = true;
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
