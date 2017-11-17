import { Component, Input } from "@angular/core";
import { Hero } from "../@backend/models";

@Component({
    selector: 'card',
    templateUrl: './card.component.html'
})
export class CardComponent {
    @Input()
    hero: Hero;
}
