import { HttpClientModule } from '@angular/common/http/src/module';
import { Component } from '@angular/core';
import { OnInit } from '@angular/core/src/metadata/lifecycle_hooks';
import { HttpClient } from '@angular/common/http';
import { Hero } from './@backend/models';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  heroes: Hero[];

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    this.http.get('api/heroes').subscribe(response => {
      this.heroes = response['data']
    })
  }
}
