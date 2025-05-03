import Location from "@/views/Location.vue";
import LocationImage from "@/views/LocationImage.vue";
import People from "@/views/People.vue";
import PeopleFace from "@/views/PeopleFace.vue";
import Search from "@/views/Search.vue";
import Setting from "@/views/Setting.vue";
import Timeline from "@/views/Timeline.vue";
import TimelineImage from "@/views/TimelineImage.vue";
import { createRouter, createWebHashHistory } from "vue-router";

export default createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            redirect: "/setting"
        },
        {
            path: "/search",
            component: Search
        },
        {
            path: "/people",
            component: People
        },
        {
            path: "/people/:cid",
            component: PeopleFace
        },
        {
            path: "/location",
            component: Location
        },
        {
            path: "/timeline",
            component: Timeline
        },
        {
            path: "/setting",
            component: Setting
        },
        {
            path: "/timeline/:year/:month",
            component: TimelineImage
        },
        {
            path: "/location/:locs/",
            component: LocationImage
        }
    ]
})