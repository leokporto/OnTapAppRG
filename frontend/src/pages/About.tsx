
export default function About() {
    return (
        <>
            <h2 className="text-3xl font-bold mb-4 text-(--color-text-muted)">About OnTapAppRG</h2>
            <p className="mb-2 text-(--color-text-muted)">
                OnTapAppRG is a solution developed based on OnTapApp repository to manage and present information about beers and their styles, providing a modern and efficient experience for users and administrators. The project uses a modern architecture based on a modular monolith, integrating backend, frontend, and cloud/containerized infrastructure.
            </p>
            <p className="mb-2 text-(--color-text-muted)">
                The project is not intended to be used as a production application.
            </p>
            <p className="mb-2 text-(--color-text-muted)">
                The project will be deployed using docker compose and docker images. The objective of this project is to learn cloud native development and deployment and the chosen stack of technologies:

                <ul className="py-2">
                    <li>- Backend: Go</li>
                    <li>- Frontend: React with TypeScript</li>
                    <li>- Database: PostgreSQL</li>
                    <li>- Containerization: Docker</li>
                </ul>
            </p>

            <h2 className="text-3xl font-bold mb-4 text-(--color-text-muted)">About Me</h2>

            <p className="mb-2 text-(--color-text-muted)">
                I'm Leonardo Porto, a software developer with 25+ years of experience in .NET solution development and architecture, working on desktop applications, web, and distributed services. I'm passionate about cloud native technologies and modern application development. Specialist in C#, ASP.NET Core, SQL Server, and knowledge in Docker, with solid experience in team management and modernization of SCADA systems.
            </p>

            <p className="mb-2 text-(--color-text-muted)">
                I'm currently exploring Go for backend development and React with TypeScript for frontend applications, aiming to expand my expertise in cloud native solutions.
            </p>

            <p className="mb-4 text-(--color-text-muted)">
                I'm open to professional connections!<br />
                
                <span className="py-2 inline-flex items-center gap-2 ml-4">
                    <a href="https://www.linkedin.com/in/leokporto/" target="_blank" rel="noopener noreferrer" className="inline-flex items-center gap-2 text-blue-600 hover:underline">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M19 0h-14c-2.761 0-5 2.239-5 5v14c0 2.761 2.239 5 5 5h14c2.762 0 5-2.239 5-5v-14c0-2.761-2.238-5-5-5zm-11 19h-3v-10h3v10zm-1.5-11.268c-.966 0-1.75-.784-1.75-1.75s.784-1.75 1.75-1.75 1.75.784 1.75 1.75-.784 1.75-1.75 1.75zm15.5 11.268h-3v-5.604c0-1.337-.025-3.063-1.868-3.063-1.868 0-2.154 1.459-2.154 2.968v5.699h-3v-10h2.881v1.367h.041c.401-.761 1.379-1.563 2.838-1.563 3.034 0 3.595 1.997 3.595 4.59v5.606z"/>
                        </svg>
                    </a>
                    <a href="https://github.com/leokporto" target="_blank" rel="noopener noreferrer" className="inline-flex items-center gap-2 text-gray-800 hover:underline">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M12 0C5.37 0 0 5.373 0 12c0 5.303 3.438 9.8 8.205 11.387.6.113.82-.258.82-.577 0-.285-.011-1.04-.017-2.04-3.338.726-4.042-1.61-4.042-1.61-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.085 1.84 1.237 1.84 1.237 1.07 1.834 2.809 1.304 3.495.997.108-.775.418-1.305.762-1.605-2.665-.304-5.466-1.334-5.466-5.931 0-1.31.469-2.381 1.236-3.221-.124-.303-.535-1.523.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.984-.399 3.003-.404 1.018.005 2.046.138 3.006.404 2.291-1.553 3.297-1.23 3.297-1.23.653 1.653.242 2.873.119 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.804 5.625-5.475 5.921.43.371.823 1.102.823 2.222 0 1.606-.015 2.898-.015 3.293 0 .321.216.694.825.576C20.565 21.796 24 17.299 24 12c0-6.627-5.373-12-12-12z"/>
                        </svg>                        
                    </a>
                </span>

            </p>

            <p className="text-(--color-text-muted)">
                Cheers to discovering your next favorite beer with OnTapAppRG! 🍺
            </p>
        </>
    );
}